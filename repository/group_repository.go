package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/mapper"
	"github.com/itsektionen/mimer/model"
)

type GroupRepository interface {
	List(ctx context.Context) ([]model.Group, error)
	GetByID(ctx context.Context, id uuid.UUID) (model.Group, error)
	Create(ctx context.Context, params db.CreateCommitteeParams) (model.Group, error)
	Update(ctx context.Context, params db.UpdateCommitteeParams) (model.Group, error)
	Delete(ctx context.Context, id uuid.UUID) (model.Group, error)
}

type groupRepository struct {
	q db.Querier
}

func NewGroupRepository(q db.Querier) GroupRepository {
	return &groupRepository{q: q}
}

func (r *groupRepository) List(ctx context.Context) ([]model.Group, error) {
	committees, err := r.q.ListCommittees(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.ToGroups(committees), nil
}

func (r *groupRepository) GetByID(ctx context.Context, id uuid.UUID) (model.Group, error) {
	committee, err := r.q.GetCommittee(ctx, id)
	if err != nil {
		return model.Group{}, err
	}
	return mapper.ToGroup(committee), nil
}

func (r *groupRepository) Create(ctx context.Context, params db.CreateCommitteeParams) (model.Group, error) {
	committee, err := r.q.CreateCommittee(ctx, params)
	if err != nil {
		return model.Group{}, err
	}
	return mapper.ToGroup(committee), nil
}

func (r *groupRepository) Update(ctx context.Context, params db.UpdateCommitteeParams) (model.Group, error) {
	committee, err := r.q.UpdateCommittee(ctx, params)
	if err != nil {
		return model.Group{}, err
	}
	return mapper.ToGroup(committee), nil
}

func (r *groupRepository) Delete(ctx context.Context, id uuid.UUID) (model.Group, error) {
	committee, err := r.q.DeleteCommittee(ctx, id)
	if err != nil {
		return model.Group{}, err
	}
	return mapper.ToGroup(committee), nil
}
