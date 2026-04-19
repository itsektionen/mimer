package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/model"
	"github.com/itsektionen/mimer/internal/repository"
)

type PositionService interface {
	GetAllPositions(ctx context.Context) ([]model.Position, error)
	CreatePosition(ctx context.Context, params db.CreatePositionParams) (model.Position, error)
	GetPositionById(ctx context.Context, id uuid.UUID) (model.Position, error)
	AssignPosition(ctx context.Context, positionID uuid.UUID, personID uuid.UUID) (*model.Trustee, error)
}

type positionService struct {
	repo        repository.PositionRepository
	trusteeRepo repository.TrusteeRepository
}

func NewPositionService(repo repository.PositionRepository, trusteeRepo repository.TrusteeRepository) PositionService {
	return &positionService{repo: repo, trusteeRepo: trusteeRepo}
}

func (s *positionService) GetAllPositions(ctx context.Context) ([]model.Position, error) {
	return s.repo.List(ctx)
}

func (s *positionService) CreatePosition(ctx context.Context, params db.CreatePositionParams) (model.Position, error) {
	return s.repo.Create(ctx, params)
}

func (s *positionService) GetPositionById(ctx context.Context, id uuid.UUID) (model.Position, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *positionService) AssignPosition(ctx context.Context, positionID uuid.UUID, personID uuid.UUID) (*model.Trustee, error) {
	now := time.Now()
	return s.trusteeRepo.CreateTrustee(ctx, positionID, personID, now, now.AddDate(1, 0, 0))
}
