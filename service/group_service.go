package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/model"
	"github.com/itsektionen/mimer/repository"
)

type GroupService interface {
	List(ctx context.Context, inactive bool) ([]model.Group, error)
	Create(ctx context.Context, params db.CreateCommitteeParams) (model.Group, error)
	GetByID(ctx context.Context, id uuid.UUID) (model.Group, error)
	ListTrustees(ctx context.Context, committeeID uuid.UUID, inactive bool) ([]model.Trustee, error)
}

type groupService struct {
	repo        repository.GroupRepository
	trusteeRepo repository.TrusteeRepository
}

func NewGroupService(repo repository.GroupRepository, trusteeRepo repository.TrusteeRepository) GroupService {
	return &groupService{
		repo:        repo,
		trusteeRepo: trusteeRepo,
	}
}

func (s *groupService) List(ctx context.Context, inactive bool) ([]model.Group, error) {
	committees, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	if inactive {
		return committees, nil
	}

	activeCommittees := []model.Group{}
	for _, committee := range committees {
		if committee.Active {
			activeCommittees = append(activeCommittees, committee)
		}
	}

	return activeCommittees, nil
}

func (s *groupService) Create(ctx context.Context, params db.CreateCommitteeParams) (model.Group, error) {
	return s.repo.Create(ctx, params)
}

func (s *groupService) GetByID(ctx context.Context, id uuid.UUID) (model.Group, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *groupService) ListTrustees(ctx context.Context, committeeID uuid.UUID, inactive bool) ([]model.Trustee, error) {
	trustees, err := s.trusteeRepo.ListByCommitteeID(ctx, committeeID)
	if err != nil {
		return nil, err
	}

	if inactive {
		return trustees, nil
	}

	// Filter for currently elected trustees
	now := time.Now()
	var current []model.Trustee
	for _, trustee := range trustees {
		if trustee.StartDate.After(now) && !trustee.EndDate.Before(now) {
			current = append(current, trustee)
		}
	}

	return current, nil
}
