package middleware

import (
	"net/http"

	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/pkg/util"
)

func AuthMiddleware(next http.Handler, queries db.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := r.Header
		authHeader := headers.Get("Authorization")

		if r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
			return
		}

		valid, err := queries.GetApiKeyByValue(r.Context(), authHeader)
		if err != nil {
			util.RespondWithJSON(w, http.StatusForbidden, "Not allowed")
			return
		}

		if valid.Active == false {
			util.RespondWithJSON(w, http.StatusForbidden, "Not allowed")
			return
		}

		next.ServeHTTP(w, r)
	})
}
