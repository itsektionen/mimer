package app

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/itsektionen/mimer/service"
)

func pathMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "current_path", r.URL.Path)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SetupAppRouter(
	groupService service.GroupService,
	userService service.UserService,
	positionService service.PositionService,
	trusteeService service.TrusteeService,
) *chi.Mux {
	router := chi.NewRouter()
	router.Use(pathMiddleware)

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
