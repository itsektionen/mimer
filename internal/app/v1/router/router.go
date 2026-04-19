package router

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/itsektionen/mimer/internal/app/v1/handler"
	"github.com/itsektionen/mimer/internal/service"
)

func SetupV1Router(
	committeeService service.CommitteeService,
	personService service.PersonService,
	positionService service.PositionService,
) http.Handler {
	router := chi.NewRouter()
	humaCfg := huma.DefaultConfig("Mimer", "0.0.0")
	humaCfg.DocsRenderer = huma.DocsRendererScalar

	api := humachi.New(router, humaCfg)

	committeeHandler := handler.NewCommitteeHandler(committeeService)
	personHandler := handler.NewPersonHandler(personService)
	positionHandler := handler.NewPositionHandler(positionService)

	router.HandleFunc("GET /people", personHandler.HandleGetAllPeople)
	router.HandleFunc("POST /people", personHandler.HandleCreatePerson)
	router.HandleFunc("GET /people/", personHandler.HandleGetPersonById)

	router.HandleFunc("GET /positions", positionHandler.HandleGetAllPositions)
	router.HandleFunc("POST /positions", positionHandler.HandleCreatePosition)
	router.HandleFunc("GET /positions/", positionHandler.HandleGetPositionById)

	huma.Get(api, "/committees", committeeHandler.HandleListCommittees)
	huma.Post(api, "/committees", committeeHandler.HandleCreateCommittee)
	huma.Get(api, "/committees/{id}", committeeHandler.HandleGetCommitteeById)

	router.HandleFunc("GET /health", handler.GetHealth)
	router.HandleFunc("GET /", handler.GetIndex)

	return router
}
