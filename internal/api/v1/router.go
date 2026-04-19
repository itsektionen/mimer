package v1

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/itsektionen/mimer/internal/api/middleware"
	"github.com/itsektionen/mimer/internal/api/v1/handler"
	"github.com/itsektionen/mimer/internal/service"
	"go.uber.org/zap"
)

func SetupV1Router(
	logger *zap.Logger,
	committeeService service.CommitteeService,
	personService service.PersonService,
	positionService service.PositionService,
	apiKeyService service.ApiKeyService,
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

	protected := huma.NewGroup(api)

	// NOTE: I do not know if passing the group instead of the root api is the idiomatic way to go about this, but it felt right
	authMiddleware := middleware.AuthMiddleware(protected, apiKeyService)
	protected.UseMiddleware(authMiddleware)

	committeeHandler := handler.NewCommitteeHandler(committeeService)
	personHandler := handler.NewPersonHandler(personService)
	positionHandler := handler.NewPositionHandler(positionService)

	huma.Register(
		api,
		huma.Operation{
			Path:    "/people",
			Summary: "List all people",
			Method:  http.MethodGet,
			Tags:    []string{"People"},
		},
		personHandler.HandleListPeople,
	)
	huma.Register(
		protected,
		huma.Operation{
			Path:          "/people",
			Summary:       "Create person",
			Method:        http.MethodPost,
			Tags:          []string{"People"},
			DefaultStatus: http.StatusCreated,
		},
		personHandler.HandleCreatePerson,
	)
	huma.Register(
		api,
		huma.Operation{
			Path:    "/people/{id}",
			Summary: "Get person",
			Method:  http.MethodGet,
			Tags:    []string{"People"},
		},
		personHandler.HandleGetPersonById,
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
			Path:          "/positions",
			Summary:       "Create position",
			Method:        http.MethodPost,
			Tags:          []string{"Positions"},
			DefaultStatus: http.StatusCreated,
		},
		positionHandler.HandleCreatePosition,
	)
	huma.Register(
		api,
		huma.Operation{
			Path:    "/positions/{id}",
			Summary: "Get position",
			Method:  http.MethodGet,
			Tags:    []string{"Positions"},
		},
		positionHandler.HandleGetPositionById,
	)

	huma.Register(
		api,
		huma.Operation{
			Path:    "/committees",
			Summary: "List all committees",
			Method:  http.MethodGet,
			Tags:    []string{"Committees"},
		},
		committeeHandler.HandleListCommittees,
	)
	huma.Register(
		api,
		huma.Operation{
			Path:          "/committees",
			Summary:       "Create new committee",
			Method:        http.MethodPost,
			Tags:          []string{"Committees"},
			DefaultStatus: http.StatusCreated,
		},
		committeeHandler.HandleCreateCommittee,
	)

	huma.Register(
		api,
		huma.Operation{
			Path:    "/committees/{id}",
			Summary: "Get committee",
			Method:  http.MethodGet,
			Tags:    []string{"Committees"},
		},
		committeeHandler.HandleGetCommitteeById,
	)
	huma.Register(
		api,
		huma.Operation{
			Path:    "/committees/{id}/trustees",
			Summary: "List committee trustees",
			Method:  http.MethodGet,
			Tags:    []string{"Committees"},
		},
		committeeHandler.HandleListCommitteeTrustees,
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
	router.HandleFunc("GET /", handler.GetIndex)

	return router
}
