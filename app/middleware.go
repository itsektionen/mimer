package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/itsektionen/mimer/app/ctxs"
	"github.com/itsektionen/mimer/service"
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

func sessionMiddleware(authService service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/static/") {
				next.ServeHTTP(w, r)
				return
			}

			cookie, err := r.Cookie("mimer_token")
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			token := cookie.Value
			if token == "" {
				next.ServeHTTP(w, r)
				return
			}

			claims, err := authService.ValidateToken(r.Context(), token)
			if err != nil {
				fmt.Println("session validation failed:", err)
				// Clear invalid cookie
				http.SetCookie(w, &http.Cookie{
					Name:     "mimer_token",
					Value:    "",
					Path:     "/",
					HttpOnly: true,
					MaxAge:   -1,
				})
				next.ServeHTTP(w, r)
				return
			}

			ctx := ctxs.UserIntoContext(r.Context(), claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func requireAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := ctxs.UserFromContext(r.Context())
		if !ok {
			nextPath := r.URL.Path
			redirectURL := fmt.Sprintf("/auth/login?next=%s", nextPath)
			http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}
