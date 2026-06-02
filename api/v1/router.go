package v1

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/itsektionen/mimer/api/middleware"
	"github.com/itsektionen/mimer/api/v1/handler"
	"github.com/itsektionen/mimer/service"
	"go.uber.org/zap"
)

func SetupV1Router(
	logger *zap.Logger,
	groupService service.GroupService,
	userService service.UserService,
	positionService service.PositionService,
) http.Handler {
	router := chi.NewRouter()
	humaCfg := huma.DefaultConfig("Mimer", "0.0.0")
	humaCfg.DocsRenderer = huma.DocsRendererScalar
	humaCfg.Servers = []*huma.Server{
		{URL: "/api/v1"},
	}

	api := humachi.New(router, humaCfg)
	loggingMiddleware := middleware.LoggingMiddleware(logger)
	api.UseMiddleware(loggingMiddleware)

	groupHandler := handler.NewGroupHandler(groupService)
	userHandler := handler.NewUserHandler(userService)
	positionHandler := handler.NewPositionHandler(positionService)

	huma.Register(
		api,
		huma.Operation{
			Path:    "/users",
			Summary: "List all users",
			Method:  http.MethodGet,
			Tags:    []string{"Users"},
		},
		userHandler.HandleListUsers,
	)
	huma.Register(
		api,
		huma.Operation{
			Path:    "/users/{id}",
			Summary: "Get user",
			Method:  http.MethodGet,
			Tags:    []string{"Users"},
		},
		userHandler.HandleGetUserByID,
	)

	huma.Register(
		api,
		huma.Operation{
			Path:    "/positions",
			Summary: "List all positions",
			Method:  http.MethodGet,
			Tags:    []string{"Positions"},
		},
		positionHandler.HandleListPositions,
	)
	huma.Register(
		api,
		huma.Operation{
			Path:    "/positions/{id}",
			Summary: "Get position",
			Method:  http.MethodGet,
			Tags:    []string{"Positions"},
		},
		positionHandler.HandleGetPositionByID,
	)

	huma.Register(
		api,
		huma.Operation{
			Path:    "/groups",
			Summary: "List all groups",
			Method:  http.MethodGet,
			Tags:    []string{"Groups"},
		},
		groupHandler.HandleListGroups,
	)

	huma.Register(
		api,
		huma.Operation{
			Path:    "/groups/{id}",
			Summary: "Get group",
			Method:  http.MethodGet,
			Tags:    []string{"Groups"},
		},
		groupHandler.HandleGetGroupById,
	)
	huma.Register(
		api,
		huma.Operation{
			Path:    "/groups/{id}/trustees",
			Summary: "List group trustees",
			Method:  http.MethodGet,
			Tags:    []string{"Groups"},
		},
		groupHandler.HandleListGroupTrustees,
	)

	huma.Register(
		api,
		huma.Operation{
			Path:    "/health",
			Summary: "Health Check",
			Method:  http.MethodGet,
		},
		handler.GetHealth,
	)

	return router
}
