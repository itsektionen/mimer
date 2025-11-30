package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/app/v1/service"
	"github.com/itsektionen/mimer/internal/pkg/db"
	"github.com/itsektionen/mimer/internal/pkg/util"
)

type CommitteeHandler struct {
	committeeService service.CommitteeService
}

func NewCommitteeHandler(s service.CommitteeService) *CommitteeHandler {
	return &CommitteeHandler{committeeService: s}
}

func (h *CommitteeHandler) HandleCreateCommittee(w http.ResponseWriter, r *http.Request) {
	var newCommittee db.CreateCommitteeParams
	err := json.NewDecoder(r.Body).Decode(&newCommittee)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	ctx := r.Context()

	committee, err := h.committeeService.CreateCommittee(ctx, newCommittee)
	if err != nil {
		log.Printf("%v", err)
		util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, committee)
}

func (h *CommitteeHandler) HandleGetAllCommittees(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	committees, err := h.committeeService.GetAllCommittees(ctx)
	if err != nil {
		log.Printf("%v", err)
		util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	util.RespondWithJSON(w, http.StatusOK, committees)
}

func (h *CommitteeHandler) HandleGetCommitteeById(w http.ResponseWriter, r *http.Request) {
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

	committee, err := h.committeeService.GetCommitteeById(ctx, id)
	if err != nil {
		log.Printf("%v", err)
		util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	}
	util.RespondWithJSON(w, http.StatusOK, committee)

}
