package handler

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/model"
	"github.com/itsektionen/mimer/internal/service"
)

type PersonHandler struct {
	personService service.PersonService
}

func NewPersonHandler(s service.PersonService) *PersonHandler {
	return &PersonHandler{personService: s}
}

type CreatePersonRequest struct {
	Body struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
}

type CreatePersonResponse struct {
	Body model.Person
}

func (h *PersonHandler) HandleCreatePerson(ctx context.Context, input *CreatePersonRequest) (*CreatePersonResponse, error) {
	person, err := h.personService.CreatePerson(ctx, db.CreatePersonParams{
		FirstName: input.Body.FirstName,
		LastName:  input.Body.LastName,
	})
	if err != nil {
		return nil, err
	}
	return &CreatePersonResponse{Body: person}, nil
}

type ListPeopleResponse struct {
	Body []model.Person
}

func (h *PersonHandler) HandleListPeople(ctx context.Context, input *struct{}) (*ListPeopleResponse, error) {
	resp := &ListPeopleResponse{}
	people, err := h.personService.GetAllPeople(ctx)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	resp.Body = people
	return resp, nil
}

type GetPersonByIdRequest struct {
	ID uuid.UUID `path:"id"`
}

type GetPersonByIdResponse struct {
	Body model.Person
}

func (h *PersonHandler) HandleGetPersonById(ctx context.Context, input *GetPersonByIdRequest) (*GetPersonByIdResponse, error) {
	resp := &GetPersonByIdResponse{}
	person, err := h.personService.GetPersonById(ctx, input.ID)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	resp.Body = person
	return resp, nil
}
