package handler

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/model"
	"github.com/itsektionen/mimer/service"
)

type PersonHandler struct {
	personService service.PersonService
}

func NewPersonHandler(s service.PersonService) *PersonHandler {
	return &PersonHandler{personService: s}
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
