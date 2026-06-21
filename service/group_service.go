package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/model"
	"github.com/itsektionen/mimer/repository"
)

type GroupService interface {
	ListAll(ctx context.Context) ([]model.Group, error)
	ListActive(ctx context.Context) ([]model.Group, error)
	Create(ctx context.Context, params db.CreateGroupParams) (model.Group, error)
	GetByID(ctx context.Context, id uuid.UUID) (model.Group, error)
	ListTrustees(ctx context.Context, groupID uuid.UUID, inactive bool) ([]model.Trustee, error)
	ListActiveWithPositions(ctx context.Context) ([]model.GroupWithPositions, error)
}

type groupService struct {
	repo         repository.GroupRepository
	trusteeRepo  repository.TrusteeRepository
	positionRepo repository.PositionRepository
}

func NewGroupService(
	repo repository.GroupRepository,
	trusteeRepo repository.TrusteeRepository,
	positionRepo repository.PositionRepository,
) GroupService {
	return &groupService{
		repo:         repo,
		trusteeRepo:  trusteeRepo,
		positionRepo: positionRepo,
	}
}

func (s *groupService) ListActive(ctx context.Context) ([]model.Group, error) {
	groups, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	activeGroups := []model.Group{}
	for _, group := range groups {
		if group.IsActive() {
			activeGroups = append(activeGroups, group)
		}
	}

	return activeGroups, nil
}

func (s *groupService) ListAll(ctx context.Context) ([]model.Group, error) {
	groups, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (s *groupService) Create(ctx context.Context, params db.CreateGroupParams) (model.Group, error) {
	return s.repo.Create(ctx, params)
}

func (s *groupService) GetByID(ctx context.Context, id uuid.UUID) (model.Group, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *groupService) ListTrustees(ctx context.Context, groupID uuid.UUID, inactive bool) ([]model.Trustee, error) {
	trustees, err := s.trusteeRepo.ListByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}

	if inactive {
		return trustees, nil
	}

	var current []model.Trustee
	for _, trustee := range trustees {
		if trustee.IsActive() {
			current = append(current, trustee)
		}
	}

	return current, nil
}

func (s *groupService) ListActiveWithPositions(ctx context.Context) ([]model.GroupWithPositions, error) {
	groups, err := s.ListActive(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]model.GroupWithPositions, 0, len(groups))
	for _, g := range groups {
		positions, err := s.positionRepo.ListByGroupIDWithActiveTrustee(ctx, g.ID)
		if err != nil {
			return nil, err
		}

		if len(positions) > 0 {
			result = append(result, model.GroupWithPositions{
				Group:     g,
				Positions: positions,
			})
		}
	}

	return result, nil
}
