package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/model"
	"github.com/itsektionen/mimer/internal/repository"
)

type CommitteeService interface {
	GetAllCommittees(ctx context.Context) ([]model.Committee, error)
	CreateCommittee(ctx context.Context, params db.CreateCommitteeParams) (model.Committee, error)
	GetCommitteeById(ctx context.Context, id uuid.UUID) (model.Committee, error)
	GetCommitteeTrustees(ctx context.Context, committeeID uuid.UUID, currentOnly bool) ([]model.Trustee, error)
}

type committeeService struct {
	repo        repository.CommitteeRepository
	trusteeRepo repository.TrusteeRepository
}

func NewCommitteeService(repo repository.CommitteeRepository, trusteeRepo repository.TrusteeRepository) CommitteeService {
	return &committeeService{
		repo:        repo,
		trusteeRepo: trusteeRepo,
	}
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

func (s *committeeService) GetCommitteeTrustees(ctx context.Context, committeeID uuid.UUID, currentOnly bool) ([]model.Trustee, error) {
	trustees, err := s.trusteeRepo.ListCommitteeTrustees(ctx, committeeID)
	if err != nil {
		return nil, err
	}

	if !currentOnly {
		return trustees, nil
	}

	// Filter for currently elected trustees
	now := time.Now()
	var current []model.Trustee
	for _, trustee := range trustees {
		if !trustee.StartDate.After(now) && !trustee.EndDate.Before(now) {
			current = append(current, trustee)
		}
	}

	return current, nil
}
