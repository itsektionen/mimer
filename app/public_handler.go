package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/itsektionen/mimer/model"
	"github.com/itsektionen/mimer/service"
	"github.com/itsektionen/mimer/templates/partials"
	"github.com/itsektionen/mimer/templates/views"
)

type PublicHandler struct {
	groupService    service.GroupService
	userService     service.UserService
	positionService service.PositionService
	trusteeService  service.TrusteeService
}

func NewPublicHandler(
	groupService service.GroupService,
	userService service.UserService,
	positionService service.PositionService,
	trusteeService service.TrusteeService,
) PublicHandler {
	return PublicHandler{
		groupService,
		userService,
		positionService,
		trusteeService,
	}
}

func (h *PublicHandler) HandleStatic(w http.ResponseWriter, r *http.Request) {
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	fs.ServeHTTP(w, r)
}

func (h *PublicHandler) HandleHome(w http.ResponseWriter, r *http.Request) {
	groups, err := h.groupService.ListActiveWithPositions(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = views.Index(groups).Render(r.Context(), w)
}

func (h *PublicHandler) HandleSearch(w http.ResponseWriter, r *http.Request) {
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

func (h *PublicHandler) HandlePositions(w http.ResponseWriter, r *http.Request) {
	positions, err := h.positionService.List(r.Context())
	if err != nil {
		return
	}
	_ = views.Positions(positions).Render(r.Context(), w)
}

func (h *PublicHandler) HandleGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := h.groupService.ListActive(r.Context())
	if err != nil {
		return
	}
	_ = views.Groups(groups).Render(r.Context(), w)
}

func (h *PublicHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.List(r.Context())
	if err != nil {
		return
	}
	_ = views.Users(users).Render(r.Context(), w)
}

func (h *PublicHandler) HandleGroupByID(w http.ResponseWriter, r *http.Request) {
	groupIDString := chi.URLParam(r, "groupID")
	groupID, err := uuid.Parse(groupIDString)
	if err != nil {
		fmt.Print(err)
		return
	}

	group, err := h.groupService.GetByID(r.Context(), groupID)
	if err != nil {
		fmt.Println(err)
		return
	}

	trustees, err := h.groupService.ListTrustees(r.Context(), groupID, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = views.Group(group, trustees).Render(r.Context(), w)
}

func (h *PublicHandler) HandlePositionByID(w http.ResponseWriter, r *http.Request) {
	positionIDString := chi.URLParam(r, "positionID")
	positionID, err := uuid.Parse(positionIDString)
	if err != nil {
		fmt.Println(err)
		return
	}

	position, err := h.positionService.GetByID(r.Context(), positionID)
	if err != nil {
		fmt.Println(err)
	}

	trustees, err := h.positionService.ListTrustees(r.Context(), positionID)
	if err != nil {
		fmt.Println(err)
	}

	_ = views.Position(position, trustees).Render(r.Context(), w)
}

func (h *PublicHandler) HandleNotFound(w http.ResponseWriter, r *http.Request) {
	_ = views.NotFound().Render(r.Context(), w)
}
