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
	router.Get("/positions", appHandler.HandlePositions)
	router.Get("/groups", appHandler.HandleGroups)
	router.Get("/users", appHandler.HandleUsers)
	router.Post("/search", appHandler.HandleSearch)
	router.NotFound(appHandler.HandleNotFound)

	return router
}
