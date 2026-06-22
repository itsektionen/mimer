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
	searchService service.SearchService,
) *chi.Mux {
	router := chi.NewRouter()
	router.Use(pathMiddleware)
	router.Use(themeMiddleware)

	publicHandler := NewPublicHandler(
		groupService,
		userService,
		positionService,
		trusteeService,
		searchService,
	)

	router.Get("/static/*", publicHandler.HandleStatic)
	router.Get("/", publicHandler.HandleHome)
	router.Get("/positions", publicHandler.HandlePositions)
	router.Get("/positions/{positionID}", publicHandler.HandlePositionByID)
	router.Get("/groups", publicHandler.HandleGroups)
	router.Get("/groups/{groupID}", publicHandler.HandleGroupByID)
	router.Get("/users", publicHandler.HandleUsers)
	router.Post("/search", publicHandler.HandleSearch)
	router.Post("/toggle-theme", publicHandler.HandleToggleTheme)
	router.NotFound(publicHandler.HandleNotFound)

	return router
}
