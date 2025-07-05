package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/itsektionen/mimer/internal/pkg/util"
	"github.com/itsektionen/mimer/internal/repository"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf(
			"%s [%s] %s\n",
			time.Now().Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL,
		)
		next.ServeHTTP(w, r)
	})
}

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
