package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	v1 "github.com/itsektionen/mimer/api/v1"
	"github.com/itsektionen/mimer/service"
	"go.uber.org/zap"
)

func SetupAPIRouter(
	logger *zap.Logger,
	committeeService service.GroupService,
	personService service.PersonService,
	positionService service.PositionService,
) http.Handler {
	router := chi.NewRouter()

	apiV1Router := v1.SetupV1Router(
		logger,
		committeeService,
		personService,
		positionService,
	)
	router.Mount("/v1", apiV1Router)

	return router
}
