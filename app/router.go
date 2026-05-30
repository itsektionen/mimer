package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/itsektionen/mimer/service"
	"github.com/itsektionen/mimer/templates"
)

func SetupAppRouter(
	committeeService service.CommitteeService,
	personService service.PersonService,
	positionService service.PositionService,
	apiKeyService service.ApiKeyService,
) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		positions, err := positionService.GetAllPositions(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		templates.Index(positions).Render(r.Context(), w)
	})

	router.Get("/static/{filename}", func(w http.ResponseWriter, r *http.Request) {
		filename := chi.URLParam(r, "filename")
		http.ServeFile(w, r, "./static/"+filename)
	})

	return router
}
