package service

import (
	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/model"
	"github.com/itsektionen/mimer/internal/repository"
)

type PersonService interface {
	GetAllPeople() ([]model.Person, error)
	CreatePerson(*model.Person) (*model.Person, error)
	GetPersonById(id uuid.UUID) (*model.Person, error)
}

type personService struct {
	personRepo repository.PersonRepository
}

func NewPersonService(personRepo repository.PersonRepository) PersonService {
	return &personService{personRepo: personRepo}
}

func (s *personService) GetAllPeople() ([]model.Person, error) {
	return s.personRepo.GetAllPeople()
}

func (s *personService) CreatePerson(person *model.Person) (*model.Person, error) {
	return s.personRepo.CreatePerson(person)
}

func (s *personService) GetPersonById(id uuid.UUID) (*model.Person, error) {
	return s.personRepo.GetPersonById(id)
}
