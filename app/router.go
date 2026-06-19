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

	publicHandler := NewPublicHandler(groupService, userService, positionService, trusteeService)

	router.Get("/static/*", publicHandler.HandleStatic)
	router.Get("/", publicHandler.HandleHome)
	router.Get("/positions", publicHandler.HandlePositions)
	router.Get("/positions/{positionID}", publicHandler.HandlePositionByID)
	router.Get("/groups", publicHandler.HandleGroups)
	router.Get("/groups/{groupID}", publicHandler.HandleGroupByID)
	router.Get("/users", publicHandler.HandleUsers)
	router.Post("/search", publicHandler.HandleSearch)
	router.NotFound(publicHandler.HandleNotFound)

	return router
}
