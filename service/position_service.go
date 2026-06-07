package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/model"
	"github.com/itsektionen/mimer/repository"
)

type PositionService interface {
	List(ctx context.Context) ([]model.Position, error)
	Create(ctx context.Context, params db.CreatePositionParams) (model.Position, error)
	GetByID(ctx context.Context, id uuid.UUID) (model.Position, error)
	Assign(ctx context.Context, positionID uuid.UUID, personID uuid.UUID) (*model.Trustee, error)
}

type positionService struct {
	repo        repository.PositionRepository
	trusteeRepo repository.TrusteeRepository
}

func NewPositionService(repo repository.PositionRepository, trusteeRepo repository.TrusteeRepository) PositionService {
	return &positionService{repo: repo, trusteeRepo: trusteeRepo}
}

func (s *positionService) List(ctx context.Context) ([]model.Position, error) {
	return s.repo.List(ctx)
}

func (s *positionService) Create(ctx context.Context, params db.CreatePositionParams) (model.Position, error) {
	return s.repo.Create(ctx, params)
}

func (s *positionService) GetByID(ctx context.Context, id uuid.UUID) (model.Position, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *positionService) Assign(ctx context.Context, positionID uuid.UUID, personID uuid.UUID) (*model.Trustee, error) {
	now := time.Now()
	return s.trusteeRepo.Create(ctx, positionID, personID, now, now.AddDate(1, 0, 0))
}
