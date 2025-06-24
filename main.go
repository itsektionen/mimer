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
	v1Router "github.com/itsektionen/mimer/internal/app/v1/router"
	v1Service "github.com/itsektionen/mimer/internal/app/v1/service"
	"github.com/itsektionen/mimer/internal/pkg/db"
	"github.com/itsektionen/mimer/internal/repository"
	"github.com/itsektionen/mimer/internal/router"
)

//go:embed migrations/*.sql
var migrations embed.FS

func main() {
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

	migrations, err := iofs.New(migrations, "migrations")
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

	committeeRepo := repository.NewCommitteeRepository(dbConn)
	committeeService := v1Service.NewCommitteeService(committeeRepo)
	personRepo := repository.NewPersonRepository(dbConn)
	personService := v1Service.NewPersonService(personRepo)

	v1APIRouter := v1Router.SetupV1Router(committeeService, personService)
	rootMux := router.SetupRootRouter(v1APIRouter)

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", rootMux))
}
