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
	Create(ctx context.Context, params db.CreateGroupParams) (model.Group, error)
	Update(ctx context.Context, params db.UpdateGroupParams) (model.Group, error)
	Delete(ctx context.Context, id uuid.UUID) (model.Group, error)
}

type groupRepository struct {
	q db.Querier
}

func NewGroupRepository(q db.Querier) GroupRepository {
	return &groupRepository{q: q}
}

func (r *groupRepository) List(ctx context.Context) ([]model.Group, error) {
	groups, err := r.q.ListGroups(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.ToGroups(groups), nil
}

func (r *groupRepository) GetByID(ctx context.Context, id uuid.UUID) (model.Group, error) {
	group, err := r.q.GetGroup(ctx, id)
	if err != nil {
		return model.Group{}, err
	}
	return mapper.ToGroup(group), nil
}

func (r *groupRepository) Create(ctx context.Context, params db.CreateGroupParams) (model.Group, error) {
	group, err := r.q.CreateGroup(ctx, params)
	if err != nil {
		return model.Group{}, err
	}
	return mapper.ToGroup(group), nil
}

func (r *groupRepository) Update(ctx context.Context, params db.UpdateGroupParams) (model.Group, error) {
	group, err := r.q.UpdateGroup(ctx, params)
	if err != nil {
		return model.Group{}, err
	}
	return mapper.ToGroup(group), nil
}

func (r *groupRepository) Delete(ctx context.Context, id uuid.UUID) (model.Group, error) {
	group, err := r.q.DeleteGroup(ctx, id)
	if err != nil {
		return model.Group{}, err
	}
	return mapper.ToGroup(group), nil
}
