package v1

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/itsektionen/mimer/internal/api/middleware"
	"github.com/itsektionen/mimer/internal/api/v1/handler"
	"github.com/itsektionen/mimer/internal/service"
)

func SetupV1Router(
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
	api.UseMiddleware(middleware.LoggingMiddleware)

	protected := huma.NewGroup(api)

	// NOTE: I do not know if passing the group instead of the root api is the idiomatic way to go about this, but it felt right
	authMiddleware := middleware.NewAuthMiddleware(protected, apiKeyService)
	protected.UseMiddleware(authMiddleware.Handle)

	committeeHandler := handler.NewCommitteeHandler(committeeService)
	personHandler := handler.NewPersonHandler(personService)
	positionHandler := handler.NewPositionHandler(positionService)

	huma.Get(api, "/people", personHandler.HandleListPeople)
	huma.Post(protected, "/people", personHandler.HandleCreatePerson)
	huma.Get(api, "/people/{id}", personHandler.HandleGetPersonById)

	huma.Get(api, "/positions", positionHandler.HandleListPositions)
	huma.Post(protected, "/positions", positionHandler.HandleCreatePosition)
	huma.Get(api, "/positions/{id}", positionHandler.HandleGetPositionById)

	huma.Get(api, "/committees", committeeHandler.HandleListCommittees)
	huma.Post(protected, "/committees", committeeHandler.HandleCreateCommittee)
	huma.Get(api, "/committees/{id}", committeeHandler.HandleGetCommitteeById)
	huma.Get(api, "/committees/{id}/trustees", committeeHandler.HandleGetCommitteeTrustees)

	huma.Get(api, "/health", handler.GetHealth)
	router.HandleFunc("GET /", handler.GetIndex)

	return router
}
