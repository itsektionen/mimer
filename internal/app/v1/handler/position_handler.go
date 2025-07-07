package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/app/v1/service"
	"github.com/itsektionen/mimer/internal/model"
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
		positions, err := h.positionService.GetAllPositions()
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

	switch r.Method {
	case http.MethodGet:
		position, err := h.positionService.GetPositionById(id)
		if err != nil {
			log.Printf("%v", err)
			util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		}
		util.RespondWithJSON(w, http.StatusOK, position)
	}
}

func (h *PositionHandler) createPosition(w http.ResponseWriter, r *http.Request) {
	var position *model.Position
	err := json.NewDecoder(r.Body).Decode(&position)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	position, err = h.positionService.CreatePosition(position)
	if err != nil {
		log.Printf("%v", err)
		util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, position)
}
