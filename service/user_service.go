package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/model"
	"github.com/itsektionen/mimer/repository"
)

type UserService interface {
	List(ctx context.Context) ([]model.User, error)
	Create(ctx context.Context, params db.CreateUserParams) (model.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) List(ctx context.Context) ([]model.User, error) {
	return s.repo.List(ctx)
}

func (s *userService) Create(ctx context.Context, params db.CreateUserParams) (model.User, error) {
	return s.repo.Create(ctx, params)
}

func (s *userService) GetByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	return s.repo.GetByID(ctx, id)
}
