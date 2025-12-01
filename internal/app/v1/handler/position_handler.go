package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/app/v1/service"
	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/response"
)

type PositionHandler struct {
	positionService service.PositionService
}

func NewPositionHandler(s service.PositionService) *PositionHandler {
	return &PositionHandler{positionService: s}
}

func (h *PositionHandler) HandleCreatePosition(w http.ResponseWriter, r *http.Request) {
	var newPosition db.CreatePositionParams
	err := json.NewDecoder(r.Body).Decode(&newPosition)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	ctx := r.Context()

	position, err := h.positionService.CreatePosition(ctx, newPosition)
	if err != nil {
		log.Printf("%v", err)
		response.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	response.RespondWithJSON(w, http.StatusCreated, position)
}

func (h *PositionHandler) HandleGetAllPositions(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	positions, err := h.positionService.GetAllPositions(ctx)
	if err != nil {
		log.Printf("%v", err)
		response.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	response.RespondWithJSON(w, http.StatusOK, positions)

}

func (h *PositionHandler) HandleGetPositionById(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 2 {
		response.RespondWithError(w, http.StatusBadRequest, "Invalid URL path")
		return
	}
	idStr := pathSegments[len(pathSegments)-1]
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, "Invalid UUID")
	}

	ctx := r.Context()

	position, err := h.positionService.GetPositionById(ctx, id)
	if err != nil {
		log.Printf("%v", err)
		response.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	}
	response.RespondWithJSON(w, http.StatusOK, position)

}
