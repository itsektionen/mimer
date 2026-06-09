package app

import (
	"net/http"
	"strings"

	"github.com/itsektionen/mimer/model"
	"github.com/itsektionen/mimer/service"
	"github.com/itsektionen/mimer/templates/partials"
	"github.com/itsektionen/mimer/templates/views"
)

type AppHandler struct {
	groupService    service.GroupService
	userService     service.UserService
	positionService service.PositionService
	trusteeService  service.TrusteeService
}

func NewAppHandler(
	groupService service.GroupService,
	userService service.UserService,
	positionService service.PositionService,
	trusteeService service.TrusteeService,
) AppHandler {
	return AppHandler{
		groupService,
		userService,
		positionService,
		trusteeService,
	}
}

func (h *AppHandler) HandleStatic(w http.ResponseWriter, r *http.Request) {
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	fs.ServeHTTP(w, r)
}

func (h *AppHandler) HandleHome(w http.ResponseWriter, r *http.Request) {
	trustees, err := h.trusteeService.ListActive(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = views.Index(trustees).Render(r.Context(), w)
}

func (h *AppHandler) HandleSearch(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}
	searchQuery := r.FormValue("search")

	if searchQuery == "" {
		return
	}

	searchResults := []model.Position{}
	positions, err := h.positionService.List(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: Implement reasonable search algorithm
	for _, position := range positions {
		if strings.Contains(strings.ToLower(position.Name), strings.ToLower(searchQuery)) {
			searchResults = append(searchResults, position)
		}
	}

	// TODO: Figure out how to clear search result
	if len(searchResults) == 0 {
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	for _, position := range searchResults {
		if err := partials.SearchResult(position.Name, "/positions/"+position.ID.String()).Render(r.Context(), w); err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
	}
}

func (h *AppHandler) HandlePositions(w http.ResponseWriter, r *http.Request) {
	positions, err := h.positionService.List(r.Context())
	if err != nil {
		return
	}
	_ = views.Positions(positions).Render(r.Context(), w)
}

func (h *AppHandler) HandleGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := h.groupService.ListActive(r.Context())
	if err != nil {
		return
	}
	_ = views.Groups(groups).Render(r.Context(), w)
}

func (h *AppHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.List(r.Context())
	if err != nil {
		return
	}
	_ = views.Users(users).Render(r.Context(), w)
}

func (h *AppHandler) HandleNotFound(w http.ResponseWriter, r *http.Request) {
	_ = views.NotFound().Render(r.Context(), w)
}
