package middleware

import (
	"net/http"

	"github.com/itsektionen/mimer/internal/pkg/util"
	"github.com/itsektionen/mimer/internal/repository"
)

func AuthMiddleware(next http.Handler, repository repository.ApiKeyRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := r.Header
		authHeader := headers.Get("Authorization")

		valid, err := repository.GetApiKeyByValue(authHeader)
		if err != nil {
			util.RespondWithJSON(w, http.StatusForbidden, "Not allowed")
			return
		}

		if valid == nil {
			util.RespondWithJSON(w, http.StatusForbidden, "Not allowed")
			return
		}

		next.ServeHTTP(w, r)
	})
}
