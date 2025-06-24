package router

import (
	"net/http"
)

func SetupRootRouter(
	v1APIRouter http.Handler,
) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1APIRouter))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/api/v1/", http.StatusMovedPermanently)
			return
		}
		http.NotFound(w, r)
	})

	return mux
}
