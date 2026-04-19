package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/model"
	"github.com/itsektionen/mimer/internal/repository"
)

type PositionService interface {
	GetAllPositions(ctx context.Context) ([]model.Position, error)
	CreatePosition(ctx context.Context, params db.CreatePositionParams) (model.Position, error)
	GetPositionById(ctx context.Context, id uuid.UUID) (model.Position, error)
}

type positionService struct {
	repo repository.PositionRepository
}

func NewPositionService(repo repository.PositionRepository) PositionService {
	return &positionService{repo: repo}
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
