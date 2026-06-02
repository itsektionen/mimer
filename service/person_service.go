package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/model"
	"github.com/itsektionen/mimer/repository"
)

type PersonService interface {
	List(ctx context.Context) ([]model.Person, error)
	Create(ctx context.Context, params db.CreatePersonParams) (model.Person, error)
	GetByID(ctx context.Context, id uuid.UUID) (model.Person, error)
}

type personService struct {
	repo repository.PersonRepository
}

func NewPersonService(repo repository.PersonRepository) PersonService {
	return &personService{repo: repo}
}

func (s *personService) List(ctx context.Context) ([]model.Person, error) {
	return s.repo.List(ctx)
}

func (s *personService) Create(ctx context.Context, params db.CreatePersonParams) (model.Person, error) {
	return s.repo.Create(ctx, params)
}

func (s *personService) GetByID(ctx context.Context, id uuid.UUID) (model.Person, error) {
	return s.repo.GetByID(ctx, id)
}
