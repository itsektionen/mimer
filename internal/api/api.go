package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	v1 "github.com/itsektionen/mimer/internal/api/v1"
	"github.com/itsektionen/mimer/internal/service"
)

func SetupRouter(
	committeeService service.CommitteeService,
	personService service.PersonService,
	positionService service.PositionService,
) http.Handler {
	router := chi.NewRouter()

	apiV1Router := v1.SetupV1Router(
		committeeService,
		personService,
		positionService,
	)
	router.Mount("/api/v1", apiV1Router)

	return router
}
