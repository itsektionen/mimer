package app

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/service"
	"github.com/itsektionen/mimer/templates/views"
	"github.com/jackc/pgx/v5/pgtype"
)

type AdminHandler struct {
	positionService service.PositionService
	groupService    service.GroupService
}

func NewAdminHandler(
	positionService service.PositionService,
	groupService service.GroupService,
) AdminHandler {
	return AdminHandler{
		positionService,
		groupService,
	}
}

func (h *AdminHandler) HandlePositions(w http.ResponseWriter, r *http.Request) {
	positions, err := h.positionService.List(r.Context())
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	_ = views.AdminPositions(positions).Render(r.Context(), w)
}

func (h *AdminHandler) HandleCreatePositionView(w http.ResponseWriter, r *http.Request) {
	// NOTE: We might want to list all groups here for creation of legacy positions.
	groups, err := h.groupService.ListActive(r.Context())
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	_ = views.CreatePosition(groups).Render(r.Context(), w)
}

func (h *AdminHandler) HandleCreatePosition(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	description := r.FormValue("description")
	groupIDString := r.FormValue("group")
	establishedAtString := r.FormValue("establishedAt")

	groupID, err := uuid.Parse(groupIDString)
	if err != nil {
		http.Error(w, "unable to parse group id", http.StatusBadRequest)
		return
	}

	establishedAt, err := time.Parse("2006-01-02", establishedAtString)
	if err != nil {
		http.Error(w, "unable to parse established at date", http.StatusBadRequest)
		return
	}

	if name == "" || email == "" {
		http.Error(w, "missing fields", http.StatusBadRequest)
		return
	}

	pos, err := h.positionService.Create(r.Context(), db.CreatePositionParams{
		Name:        name,
		Email:       email,
		Description: &description,
		EstablishedAt: pgtype.Date{
			Time:  establishedAt,
			Valid: true,
		},
		GroupID: groupID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HXRedirect(w, "/admin/positions/"+pos.ID.String())
}

func (h *AdminHandler) HandleEditPositionView(w http.ResponseWriter, r *http.Request) {
	positionIDString := r.PathValue("positionID")
	positionID, err := uuid.Parse(positionIDString)
	if err != nil {
		http.Error(w, "unable to parse position id", http.StatusBadRequest)
		return
	}

	pos, err := h.positionService.GetByID(r.Context(), positionID)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	groups, err := h.groupService.ListActive(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = views.EditPositionView(pos, groups).Render(r.Context(), w)
}

func (h *AdminHandler) HandleEditPosition(w http.ResponseWriter, r *http.Request) {
	positionIDString := r.PathValue("positionID")
	positionID, err := uuid.Parse(positionIDString)
	if err != nil {
		http.Error(w, "unable to parse position id", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	description := r.FormValue("description")
	groupIDString := r.FormValue("group")
	establishedAtString := r.FormValue("establishedAt")

	groupID, err := uuid.Parse(groupIDString)
	if err != nil {
		http.Error(w, "unable to parse group id", http.StatusBadRequest)
		return
	}

	establishedAt, err := time.Parse("2006-01-02", establishedAtString)
	if err != nil {
		http.Error(w, "unable to parse established at date", http.StatusBadRequest)
		return
	}

	if name == "" || email == "" {
		http.Error(w, "missing fields", http.StatusBadRequest)
		return
	}

	pos, err := h.positionService.Update(r.Context(), positionID, db.UpdatePositionParams{
		Name:        name,
		Email:       email,
		Description: &description,
		EstablishedAt: pgtype.Date{
			Time:  establishedAt,
			Valid: true,
		},
		GroupID: groupID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HXRedirect(w, "/admin/positions/"+pos.ID.String())
}
