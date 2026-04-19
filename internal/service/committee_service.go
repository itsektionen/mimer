package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/model"
	"github.com/itsektionen/mimer/internal/repository"
)

type CommitteeService interface {
	GetAllCommittees(ctx context.Context) ([]model.Committee, error)
	CreateCommittee(ctx context.Context, params db.CreateCommitteeParams) (model.Committee, error)
	GetCommitteeById(ctx context.Context, id uuid.UUID) (model.Committee, error)
}

type committeeService struct {
	repo repository.CommitteeRepository
}

func NewCommitteeService(repo repository.CommitteeRepository) CommitteeService {
	return &committeeService{repo: repo}
}

func (s *committeeService) GetAllCommittees(ctx context.Context) ([]model.Committee, error) {
	return s.repo.List(ctx)
}

func (s *committeeService) CreateCommittee(ctx context.Context, params db.CreateCommitteeParams) (model.Committee, error) {
	return s.repo.Create(ctx, params)
}

func (s *committeeService) GetCommitteeById(ctx context.Context, id uuid.UUID) (model.Committee, error) {
	return s.repo.GetByID(ctx, id)
}
