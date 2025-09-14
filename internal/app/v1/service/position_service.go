package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/pkg/db"
)

type PositionService interface {
	GetAllPositions(ctx context.Context) ([]db.Position, error)
	CreatePosition(ctx context.Context, position db.CreatePositionParams) (db.Position, error)
	GetPositionById(ctx context.Context, id uuid.UUID) (db.Position, error)
}

func NewPositionService(queries db.Queries) PositionService {
	return &positionService{queries: queries}
}

type positionService struct {
	queries db.Queries
}

func (s *positionService) GetAllPositions(ctx context.Context) ([]db.Position, error) {
	return s.queries.ListPositions(ctx)
}

func (s *positionService) CreatePosition(ctx context.Context, position db.CreatePositionParams) (db.Position, error) {
	return s.queries.CreatePosition(ctx, position)
}

func (s *positionService) GetPositionById(ctx context.Context, id uuid.UUID) (db.Position, error) {
	return s.queries.GetPosition(ctx, id)
}
