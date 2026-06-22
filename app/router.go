package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/itsektionen/mimer/service"
)

func SetupAppRouter(
	groupService service.GroupService,
	userService service.UserService,
	positionService service.PositionService,
	trusteeService service.TrusteeService,
	searchService service.SearchService,
	authService service.AuthService,
) *chi.Mux {
	router := chi.NewRouter()

	publicRouter := chi.NewRouter()
	protectedRouter := chi.NewRouter()

	publicRouter.Use(pathMiddleware)
	publicRouter.Use(themeMiddleware)
	protectedRouter.Use(authMiddleware)

	publicHandler := NewPublicHandler(
		groupService,
		userService,
		positionService,
		trusteeService,
		searchService,
	)

	authHandler := NewAuthHandler(authService)

	router.Mount("/", publicRouter)
	publicRouter.Get("/static/*", publicHandler.HandleStatic)
	publicRouter.Get("/", publicHandler.HandleHome)
	publicRouter.Get("/positions", publicHandler.HandlePositions)
	publicRouter.Get("/positions/{positionID}", publicHandler.HandlePositionByID)
	publicRouter.Get("/groups", publicHandler.HandleGroups)
	publicRouter.Get("/groups/{groupID}", publicHandler.HandleGroupByID)
	publicRouter.Get("/users", publicHandler.HandleUsers)
	publicRouter.Post("/search", publicHandler.HandleSearch)
	publicRouter.Post("/toggle-theme", publicHandler.HandleToggleTheme)
	publicRouter.NotFound(publicHandler.HandleNotFound)

	publicRouter.Get("/auth/login", authHandler.Login)
	publicRouter.Get("/auth/authentik/callback", authHandler.HandleAuthentikCallback)

	router.Mount("/admin/{$}", protectedRouter)
	protectedRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("yello"))
	})

	return router
}
