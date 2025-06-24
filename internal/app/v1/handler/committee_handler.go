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

type CommitteeHandler struct {
	committeeService service.CommitteeService
}

func NewCommitteeHandler(s service.CommitteeService) *CommitteeHandler {
	return &CommitteeHandler{committeeService: s}
}

func (h *CommitteeHandler) HandleCommittees(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createCommittee(w, r)
	case http.MethodGet:
		committees, err := h.committeeService.GetAllCommittees()
		if err != nil {
			log.Printf("%v", err)
			util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		util.RespondWithJSON(w, http.StatusOK, committees)
	}

}

func (h *CommitteeHandler) HandleCommitteeById(w http.ResponseWriter, r *http.Request) {
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
		committee, err := h.committeeService.GetCommitteeById(id)
		if err != nil {
			log.Printf("%v", err)
			util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		}
		util.RespondWithJSON(w, http.StatusOK, committee)
	}
}

func (h *CommitteeHandler) createCommittee(w http.ResponseWriter, r *http.Request) {
	var committee *model.Committee
	err := json.NewDecoder(r.Body).Decode(&committee)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	committee, err = h.committeeService.CreateCommittee(committee)
	if err != nil {
		log.Printf("%v", err)
		util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, committee)
}
