package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/itsektionen/mimer/model"
	"github.com/itsektionen/mimer/service"
	"github.com/itsektionen/mimer/templates"
)

func SetupAppRouter(
	committeeService service.CommitteeService,
	personService service.PersonService,
	positionService service.PositionService,
	trusteeService service.TrusteeService,
) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		trustees, err := trusteeService.ListActive(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		templates.Index(trustees).Render(r.Context(), w)
	})

	router.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
		fs.ServeHTTP(w, r)
	})

	router.Post("/search", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "invalid form data", http.StatusBadRequest)
			return
		}
		searchQuery := r.FormValue("search")

		if searchQuery == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		searchResults := []model.Position{}
		positions, err := positionService.GetAllPositions(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		for _, position := range positions {
			if strings.HasPrefix(strings.ToLower(position.Name), strings.ToLower(searchQuery)) {
				searchResults = append(searchResults, position)
			}
		}

		if len(searchResults) == 0 {
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		resp := fmt.Sprintf(`
		<a href="/positions/%s">%s</a>
		`,
			searchResults[0].ID,
			searchResults[0].Name,
		)

		fmt.Fprint(w, resp)
	})

	return router
}
