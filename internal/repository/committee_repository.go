package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/mapper"
	"github.com/itsektionen/mimer/internal/model"
)

type CommitteeRepository interface {
	List(ctx context.Context) ([]model.Committee, error)
	GetByID(ctx context.Context, id uuid.UUID) (model.Committee, error)
	Create(ctx context.Context, params db.CreateCommitteeParams) (model.Committee, error)
	Update(ctx context.Context, params db.UpdateCommitteeParams) (model.Committee, error)
	Delete(ctx context.Context, id uuid.UUID) (model.Committee, error)
}

type committeeRepository struct {
	q db.Querier
}

func NewCommitteeRepository(q db.Querier) CommitteeRepository {
	return &committeeRepository{q: q}
}

func (r *committeeRepository) List(ctx context.Context) ([]model.Committee, error) {
	committees, err := r.q.ListCommittees(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.CommitteesFromDB(committees), nil
}

func (r *committeeRepository) GetByID(ctx context.Context, id uuid.UUID) (model.Committee, error) {
	committee, err := r.q.GetCommittee(ctx, id)
	if err != nil {
		return model.Committee{}, err
	}
	return mapper.CommitteeFromDB(committee), nil
}

func (r *committeeRepository) Create(ctx context.Context, params db.CreateCommitteeParams) (model.Committee, error) {
	committee, err := r.q.CreateCommittee(ctx, params)
	if err != nil {
		return model.Committee{}, err
	}
	return mapper.CommitteeFromDB(committee), nil
}

func (r *committeeRepository) Update(ctx context.Context, params db.UpdateCommitteeParams) (model.Committee, error) {
	committee, err := r.q.UpdateCommittee(ctx, params)
	if err != nil {
		return model.Committee{}, err
	}
	return mapper.CommitteeFromDB(committee), nil
}

func (r *committeeRepository) Delete(ctx context.Context, id uuid.UUID) (model.Committee, error) {
	committee, err := r.q.DeleteCommittee(ctx, id)
	if err != nil {
		return model.Committee{}, err
	}
	return mapper.CommitteeFromDB(committee), nil
}
