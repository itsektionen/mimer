package app

import (
	"context"
	"net/http"

	"github.com/itsektionen/mimer/app/ctxs"
)

func pathMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ctxs.PathKey, r.URL.Path)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func themeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		theme := "light"
		cookie, err := r.Cookie("mimer-theme")
		if err == nil && cookie.Value == "dark" {
			theme = "dark"
		}
		ctx := context.WithValue(r.Context(), ctxs.ThemeKey, theme)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
