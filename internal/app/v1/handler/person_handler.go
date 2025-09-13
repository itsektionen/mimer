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

type PersonHandler struct {
	personService service.PersonService
}

func NewPersonHandler(s service.PersonService) *PersonHandler {
	return &PersonHandler{personService: s}
}

func (h *PersonHandler) HandlePeople(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createPerson(w, r)
	case http.MethodGet:
		ctx := r.Context()
		people, err := h.personService.GetAllPeople(ctx)
		if err != nil {
			log.Printf("%v", err)
			util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		util.RespondWithJSON(w, http.StatusOK, people)
	}

}

func (h *PersonHandler) HandlePersonById(w http.ResponseWriter, r *http.Request) {
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
		person, err := h.personService.GetPersonById(ctx, id)
		if err != nil {
			log.Printf("%v", err)
			util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		}
		util.RespondWithJSON(w, http.StatusOK, person)
	}
}

func (h *PersonHandler) createPerson(w http.ResponseWriter, r *http.Request) {
	var newPerson db.CreatePersonParams
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	ctx := r.Context()

	person, err := h.personService.CreatePerson(ctx, newPerson)
	if err != nil {
		log.Printf("%v", err)
		util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, person)
}
