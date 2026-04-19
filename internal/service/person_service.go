package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/model"
	"github.com/itsektionen/mimer/internal/repository"
)

type PersonService interface {
	GetAllPeople(ctx context.Context) ([]model.Person, error)
	CreatePerson(ctx context.Context, params db.CreatePersonParams) (model.Person, error)
	GetPersonById(ctx context.Context, id uuid.UUID) (model.Person, error)
}

type personService struct {
	repo repository.PersonRepository
}

func NewPersonService(repo repository.PersonRepository) PersonService {
	return &personService{repo: repo}
}

func (s *personService) GetAllPeople(ctx context.Context) ([]model.Person, error) {
	return s.repo.List(ctx)
}

func (s *personService) CreatePerson(ctx context.Context, params db.CreatePersonParams) (model.Person, error) {
	return s.repo.Create(ctx, params)
}

func (s *personService) GetPersonById(ctx context.Context, id uuid.UUID) (model.Person, error) {
	return s.repo.GetByID(ctx, id)
}
