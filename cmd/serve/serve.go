package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.uber.org/zap"

	"github.com/itsektionen/mimer/api"
	"github.com/itsektionen/mimer/app"
	"github.com/itsektionen/mimer/cfg"
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/repository"
	"github.com/itsektionen/mimer/service"
)

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

	queries := db.New(conn)

	committeeRepo := repository.NewCommitteeRepository(queries)
	personRepo := repository.NewPersonRepository(queries)
	positionRepo := repository.NewPositionRepository(queries)
	trusteeRepo := repository.NewTrusteeRepository(queries)

	committeeService := service.NewCommitteeService(committeeRepo, trusteeRepo)
	personService := service.NewPersonService(personRepo)
	positionService := service.NewPositionService(positionRepo, trusteeRepo)
	trusteeService := service.NewTrusteeService(trusteeRepo)

	router := app.SetupAppRouter(
		committeeService,
		personService,
		positionService,
		trusteeService,
	)
	apiRouter := api.SetupAPIRouter(
		logger,
		committeeService,
		personService,
		positionService,
	)

	router.Mount("/api", apiRouter)

	fmt.Printf("Starting server on port %d\n", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), router))
}
