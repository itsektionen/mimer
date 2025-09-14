package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/joho/godotenv"

	"github.com/itsektionen/mimer/internal/app/v1/middleware"
	v1Router "github.com/itsektionen/mimer/internal/app/v1/router"
	v1Service "github.com/itsektionen/mimer/internal/app/v1/service"
	"github.com/itsektionen/mimer/internal/pkg/db"
	sqlc "github.com/itsektionen/mimer/internal/pkg/db"
	"github.com/itsektionen/mimer/internal/router"
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
	env := os.Getenv("MIMER_ENV")
	if env == "" {
		env = "development"
	}
	fmt.Println("env:", env)

	loadedFiles, err := loadEnv(env)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	} else {
		fmt.Println("Successfully loaded env files:")
		for _, f := range loadedFiles {
			fmt.Println(" -", f)
		}
	}

	connString := os.Getenv("DATABASE_URL")
	dbConn, err := db.SetupPostgresDB(connString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize migrations")
	}

	migrations, err := iofs.New(migrations, "db/migrations")
	if err != nil {
		log.Fatalf("Failed to read migrations")
	}

	migrator, err := migrate.NewWithInstance("iofs", migrations, "mimer", driver)
	if err != nil {
		panic(fmt.Errorf("Failed to migrate 3: %v", err))
	}

	if err := migrator.Up(); err != migrate.ErrNoChange && err != nil {
		panic(fmt.Errorf("Failed to migrate 4: %v", err))
	}

	queries := sqlc.New(dbConn)

	committeeService := v1Service.NewCommitteeService(*queries)
	personService := v1Service.NewPersonService(*queries)
	positionService := v1Service.NewPositionService(*queries)

	v1APIRouter := v1Router.SetupV1Router(committeeService, personService, positionService)
	rootMux := router.SetupRootRouter(middleware.AuthMiddleware(v1APIRouter, *queries))

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", rootMux))
}
