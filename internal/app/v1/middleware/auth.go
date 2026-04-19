package middleware

import (
	"net/http"

	"github.com/itsektionen/mimer/internal/repository"
	"github.com/itsektionen/mimer/internal/response"
)

func AuthMiddleware(next http.Handler, apiKeyRepo repository.ApiKeyRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := r.Header
		authHeader := headers.Get("Authorization")

		if r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
			return
		}

		apiKey, err := apiKeyRepo.GetByValue(r.Context(), authHeader)
		if err != nil {
			response.RespondWithJSON(w, http.StatusForbidden, "Not allowed")
			return
		}

		if !apiKey.Active {
			response.RespondWithJSON(w, http.StatusForbidden, "Not allowed")
			return
		}

		next.ServeHTTP(w, r)
	})
}
