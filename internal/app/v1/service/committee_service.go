package service

import (
	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/model"
	"github.com/itsektionen/mimer/internal/repository"
)

type CommitteeService interface {
	GetAllCommittees() ([]model.Committee, error)
	CreateCommittee(*model.Committee) (*model.Committee, error)
	GetCommitteeById(id uuid.UUID) (*model.Committee, error)
}

type committeeService struct {
	committeeRepo repository.CommitteeRepository
}

func NewCommitteeService(committeeRepo repository.CommitteeRepository) CommitteeService {
	return &committeeService{committeeRepo: committeeRepo}
}

func (s *committeeService) GetAllCommittees() ([]model.Committee, error) {
	return s.committeeRepo.GetAllCommittees()
}

func (s *committeeService) CreateCommittee(committee *model.Committee) (*model.Committee, error) {
	return s.committeeRepo.CreateCommittee(committee)
}

func (s *committeeService) GetCommitteeById(id uuid.UUID) (*model.Committee, error) {
	return s.committeeRepo.GetCommitteeById(id)
}
