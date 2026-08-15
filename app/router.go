package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/itsektionen/mimer/app/ctxs"
	"github.com/itsektionen/mimer/service"
	"github.com/itsektionen/mimer/templates/views"
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

	router.Use(sessionMiddleware(authService))
	router.Use(pathMiddleware)
	router.Use(themeMiddleware)

	publicRouter := chi.NewRouter()
	protectedRouter := chi.NewRouter()

	protectedRouter.Use(requireAuthMiddleware)

	publicHandler := NewPublicHandler(
		groupService,
		userService,
		positionService,
		trusteeService,
		searchService,
	)
	adminHandler := NewAdminHandler(
		positionService,
		groupService,
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
	router.NotFound(publicHandler.HandleNotFound)

	publicRouter.Get("/auth/login", authHandler.Login)
	publicRouter.Get("/auth/authentik/callback", authHandler.HandleAuthentikCallback)
	publicRouter.Get("/auth/logout", authHandler.Logout)

	router.Mount("/admin", protectedRouter)
	protectedRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		user, ok := ctxs.UserFromContext(r.Context())
		if !ok {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}
		_ = views.Admin(*user).Render(r.Context(), w)
	})
	protectedRouter.Get("/positions", adminHandler.HandlePositions)
	protectedRouter.Get("/positions/create", adminHandler.HandleCreatePositionView)
	protectedRouter.Post("/positions/create", adminHandler.HandleCreatePosition)

	return router
}
