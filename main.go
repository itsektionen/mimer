package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	pgxMigrate "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/itsektionen/mimer/cfg"
	"github.com/itsektionen/mimer/internal/api"
	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/repository"
	"github.com/itsektionen/mimer/internal/service"
)

//go:embed db/migrations/*.sql
var migrations embed.FS

func loadEnv(env string) ([]string, error) {
	files := []string{
		".env." + env + ".local",
	}
	if env != "test" {
		files = append(files, ".env.local")
	}
	files = append(files, ".env."+env, ".env")

	var loadedFiles []string
	var failedFiles []string

	for _, file := range files {
		if err := godotenv.Load(file); err == nil {
			loadedFiles = append(loadedFiles, file)
		} else {
			failedFiles = append(failedFiles, file)
		}
	}

	if len(loadedFiles) == 0 {
		return nil, fmt.Errorf("No environment variables found")
	}

	return loadedFiles, nil
}

func main() {
	cfg := cfg.Load()
	connString := cfg.Database.URL
	ctx := context.Background()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	conn, err := db.SetupPostgresDB(ctx, connString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close(ctx)

	// Create *sql.DB from pgx connection for migrations
	sqlDB := stdlib.OpenDB(*conn.Config())
	defer sqlDB.Close()

	driver, err := pgxMigrate.WithInstance(sqlDB, &pgxMigrate.Config{})
	if err != nil {
		log.Fatalf("failed to initialize migrations")
	}

	migrations, err := iofs.New(migrations, "db/migrations")
	if err != nil {
		log.Fatalf("failed to read migrations")
	}

	migrator, err := migrate.NewWithInstance("iofs", migrations, "mimer", driver)
	if err != nil {
		panic(fmt.Errorf("failed to migrate 3: %v", err))
	}

	if err := migrator.Up(); err != migrate.ErrNoChange && err != nil {
		panic(fmt.Errorf("failed to migrate 4: %v", err))
	}

	queries := db.New(conn)

	committeeRepo := repository.NewCommitteeRepository(queries)
	personRepo := repository.NewPersonRepository(queries)
	positionRepo := repository.NewPositionRepository(queries)
	apiKeyRepo := repository.NewApiKeyRepository(queries)
	trusteeRepo := repository.NewTrusteeRepository(queries)

	committeeService := service.NewCommitteeService(committeeRepo, trusteeRepo)
	personService := service.NewPersonService(personRepo)
	positionService := service.NewPositionService(positionRepo, trusteeRepo)
	apiKeyService := service.NewApiKeyService(apiKeyRepo)

	router := api.SetupRouter(
		logger,
		committeeService,
		personService,
		positionService,
		apiKeyService,
	)

	fmt.Printf("Starting server on port %d\n", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), router))
}
