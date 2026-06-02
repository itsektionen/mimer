package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/mapper"
	"github.com/itsektionen/mimer/model"
)

type UserRepository interface {
	List(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (model.User, error)
	Create(ctx context.Context, params db.CreateUserParams) (model.User, error)
	Update(ctx context.Context, params db.UpdateUserParams) (model.User, error)
	Delete(ctx context.Context, id uuid.UUID) (model.User, error)
}

type userRepository struct {
	q db.Querier
}

func NewUserRepository(q db.Querier) UserRepository {
	return &userRepository{q: q}
}

func (r *userRepository) List(ctx context.Context) ([]model.User, error) {
	users, err := r.q.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.ToUsers(users), nil
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	person, err := r.q.GetUser(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	return mapper.ToUser(person), nil
}

func (r *userRepository) Create(ctx context.Context, params db.CreateUserParams) (model.User, error) {
	person, err := r.q.CreateUser(ctx, params)
	if err != nil {
		return model.User{}, err
	}
	return mapper.ToUser(person), nil
}

func (r *userRepository) Update(ctx context.Context, params db.UpdateUserParams) (model.User, error) {
	person, err := r.q.UpdateUser(ctx, params)
	if err != nil {
		return model.User{}, err
	}
	return mapper.ToUser(person), nil
}

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) (model.User, error) {
	person, err := r.q.DeleteUser(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	return mapper.ToUser(person), nil
}
