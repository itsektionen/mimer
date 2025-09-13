package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/app/v1/service"
	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/pkg/util"
)

type PositionHandler struct {
	positionService service.PositionService
}

func NewPositionHandler(s service.PositionService) *PositionHandler {
	return &PositionHandler{positionService: s}
}

func (h *PositionHandler) HandlePositions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createPosition(w, r)
	case http.MethodGet:
		ctx := r.Context()
		positions, err := h.positionService.GetAllPositions(ctx)
		if err != nil {
			log.Printf("%v", err)
			util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		util.RespondWithJSON(w, http.StatusOK, positions)
	}
}

func (h *PositionHandler) HandlePositionById(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 2 {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid URL path")
		return
	}
	idStr := pathSegments[len(pathSegments)-1]
	id, err := uuid.Parse(idStr)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid UUID")
	}

	ctx := r.Context()

	switch r.Method {
	case http.MethodGet:
		position, err := h.positionService.GetPositionById(ctx, id)
		if err != nil {
			log.Printf("%v", err)
			util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		}
		util.RespondWithJSON(w, http.StatusOK, position)
	}
}

func (h *PositionHandler) createPosition(w http.ResponseWriter, r *http.Request) {
	var newPosition db.CreatePositionParams
	err := json.NewDecoder(r.Body).Decode(&newPosition)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	ctx := r.Context()

	position, err := h.positionService.CreatePosition(ctx, newPosition)
	if err != nil {
		log.Printf("%v", err)
		util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, position)
}
