package app

import (
	"fmt"
	"net/http"

	"github.com/itsektionen/mimer/app/ctxs"
)

func pathMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := ctxs.PathIntoContext(r.Context(), r.URL.Path)
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
		ctx := ctxs.ThemeIntoContext(r.Context(), theme)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextPath := r.URL.Path

		redirectURL := fmt.Sprintf("/auth/login?next=%s", nextPath)

		cookie, err := r.Cookie("mimer_token")
		if err != nil {
			fmt.Print(err)
			http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
			return
		}

		token := cookie.Value

		if token == "" {
			fmt.Println("empty token")
			http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
			return
		}

		// TODO: Validate token (introspection)
		next.ServeHTTP(w, r)
	})
}
