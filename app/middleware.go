package app

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/itsektionen/mimer/app/ctxs"
	"github.com/itsektionen/mimer/model"
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

			var accessToken string
			accessCookie, err := r.Cookie("mimer_token")
			if err == nil {
				accessToken = accessCookie.Value
			}

			var claims *model.UserClaims
			var validateErr error

			if accessToken != "" {
				claims, validateErr = authService.ValidateToken(r.Context(), accessToken)
			}

			if validateErr == nil && claims != nil {
				ctx := ctxs.UserIntoContext(r.Context(), claims)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			refreshCookie, err := r.Cookie("mimer_refresh_token")
			if err == nil && refreshCookie.Value != "" {
				newToken, refreshErr := authService.RefreshToken(r.Context(), refreshCookie.Value)
				if refreshErr == nil && newToken != nil {
					newClaims, validateErr := authService.ValidateToken(r.Context(), newToken.AccessToken)
					if validateErr == nil && newClaims != nil {
						http.SetCookie(w, &http.Cookie{
							Name:     "mimer_token",
							Value:    newToken.AccessToken,
							Path:     "/",
							HttpOnly: true,
							Secure:   true,
							SameSite: http.SameSiteLaxMode,
							MaxAge:   int(time.Until(newToken.Expiry).Seconds()),
						})

						if newToken.RefreshToken != "" {
							http.SetCookie(w, &http.Cookie{
								Name:     "mimer_refresh_token",
								Value:    newToken.RefreshToken,
								Path:     "/",
								HttpOnly: true,
								Secure:   true,
								SameSite: http.SameSiteLaxMode,
								MaxAge:   30 * 24 * 3600, // 30 days
							})
						}

						ctx := ctxs.UserIntoContext(r.Context(), newClaims)
						next.ServeHTTP(w, r.WithContext(ctx))
						return
					}
				}
			}

			if accessToken != "" {
				http.SetCookie(w, &http.Cookie{
					Name:     "mimer_token",
					Value:    "",
					Path:     "/",
					HttpOnly: true,
					MaxAge:   -1,
				})
			}
			if err == nil && refreshCookie.Value != "" {
				http.SetCookie(w, &http.Cookie{
					Name:     "mimer_refresh_token",
					Value:    "",
					Path:     "/",
					HttpOnly: true,
					MaxAge:   -1,
				})
			}

			next.ServeHTTP(w, r)
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
