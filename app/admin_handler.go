package app

import (
	"net/http"

	"github.com/itsektionen/mimer/service"
	"github.com/itsektionen/mimer/templates/views"
)

type AdminHandler struct {
	positionService service.PositionService
}

func NewAdminHandler(positionService service.PositionService) AdminHandler {
	return AdminHandler{
		positionService,
	}
}

func (h *AdminHandler) HandlePositions(w http.ResponseWriter, r *http.Request) {
	positions, err := h.positionService.List(r.Context())
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	_ = views.AdminPositions(positions).Render(r.Context(), w)
}
