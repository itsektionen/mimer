package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.uber.org/zap"

	"github.com/itsektionen/mimer/api"
	"github.com/itsektionen/mimer/app"
	"github.com/itsektionen/mimer/config"
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/repository"
	"github.com/itsektionen/mimer/service"
)

func main() {
	cfg := config.Load()
	connString := cfg.Database.URL
	ctx := context.Background()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	pool, err := db.SetupPostgresPool(ctx, connString)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer pool.Close()

	queries := db.New(pool)

	groupRepo := repository.NewGroupRepository(queries)
	userRepo := repository.NewUserRepository(queries)
	positionRepo := repository.NewPositionRepository(queries)
	trusteeRepo := repository.NewTrusteeRepository(queries)

	groupService := service.NewGroupService(groupRepo, trusteeRepo, positionRepo)
	userService := service.NewUserService(userRepo)
	positionService := service.NewPositionService(positionRepo, trusteeRepo)
	trusteeService := service.NewTrusteeService(trusteeRepo)
	searchService := service.NewSearchService(groupRepo, positionRepo)
	authService := service.NewAuthService(cfg.OAuth.ToOAuth2Config(), cfg.OAuth.UserinfoURL)

	router := app.SetupAppRouter(
		groupService,
		userService,
		positionService,
		trusteeService,
		searchService,
		authService,
	)
	apiRouter := api.SetupAPIRouter(
		logger,
		groupService,
		userService,
		positionService,
	)

	router.Mount("/api", apiRouter)

	fmt.Printf("Starting server on port %d\n", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), router))
}
