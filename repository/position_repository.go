package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/mapper"
	"github.com/itsektionen/mimer/model"
)

type PositionRepository interface {
	List(ctx context.Context) ([]model.Position, error)
	GetByID(ctx context.Context, id uuid.UUID) (model.Position, error)
	Create(ctx context.Context, params db.CreatePositionParams) (model.Position, error)
	Update(ctx context.Context, params db.UpdatePositionParams) (model.Position, error)
	Delete(ctx context.Context, id uuid.UUID) (model.Position, error)
	ListByGroupIDWithActiveTrustee(ctx context.Context, groupID uuid.UUID) ([]model.PositionWithTrustee, error)
}

type positionRepository struct {
	q db.Querier
}

func NewPositionRepository(q db.Querier) PositionRepository {
	return &positionRepository{q: q}
}

func (r *positionRepository) List(ctx context.Context) ([]model.Position, error) {
	positions, err := r.q.ListPositions(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.ToPositions(positions), nil
}

func (r *positionRepository) GetByID(ctx context.Context, id uuid.UUID) (model.Position, error) {
	position, err := r.q.GetPosition(ctx, id)
	if err != nil {
		return model.Position{}, err
	}
	return mapper.ToPosition(position), nil
}

func (r *positionRepository) Create(ctx context.Context, params db.CreatePositionParams) (model.Position, error) {
	position, err := r.q.CreatePosition(ctx, params)
	if err != nil {
		return model.Position{}, err
	}
	return mapper.ToPosition(position), nil
}

func (r *positionRepository) Update(ctx context.Context, params db.UpdatePositionParams) (model.Position, error) {
	position, err := r.q.UpdatePosition(ctx, params)
	if err != nil {
		return model.Position{}, err
	}
	return mapper.ToPosition(position), nil
}

func (r *positionRepository) Delete(ctx context.Context, id uuid.UUID) (model.Position, error) {
	position, err := r.q.DeletePosition(ctx, id)
	if err != nil {
		return model.Position{}, err
	}
	return mapper.ToPosition(position), nil
}

func (r *positionRepository) ListByGroupIDWithActiveTrustee(ctx context.Context, groupID uuid.UUID) ([]model.PositionWithTrustee, error) {
	rows, err := r.q.ListGroupPositionsWithActiveTrustees(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return mapper.ToPositionWithTrustees(rows), nil
}
