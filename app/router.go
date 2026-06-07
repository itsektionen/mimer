package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/itsektionen/mimer/service"
)

func SetupAppRouter(
	groupService service.GroupService,
	userService service.UserService,
	positionService service.PositionService,
	trusteeService service.TrusteeService,
) *chi.Mux {
	router := chi.NewRouter()

	appHandler := NewAppHandler(groupService, userService, positionService, trusteeService)

	router.Get("/static/*", appHandler.HandleStatic)
	router.Get("/", appHandler.HandleHome)
	router.Post("/search", appHandler.HandleSearch)

	return router
}
