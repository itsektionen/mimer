package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/db"
)

type PersonService interface {
	GetAllPeople(ctx context.Context) ([]db.Person, error)
	CreatePerson(ctx context.Context, person db.CreatePersonParams) (db.Person, error)
	GetPersonById(ctx context.Context, id uuid.UUID) (db.Person, error)
}

type personService struct {
	queries db.Queries
}

func NewPersonService(queries db.Queries) PersonService {
	return &personService{queries: queries}
}

func (s *personService) GetAllPeople(ctx context.Context) ([]db.Person, error) {
	return s.queries.ListPeople(ctx)
}

func (s *personService) CreatePerson(ctx context.Context, person db.CreatePersonParams) (db.Person, error) {
	return s.queries.CreatePerson(ctx, person)
}

func (s *personService) GetPersonById(ctx context.Context, id uuid.UUID) (db.Person, error) {
	return s.queries.GetPerson(ctx, id)
}
