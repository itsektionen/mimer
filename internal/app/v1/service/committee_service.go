package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/db"
)

type CommitteeService interface {
	GetAllCommittees(ctx context.Context) ([]db.Committee, error)
	CreateCommittee(ctx context.Context, committee db.CreateCommitteeParams) (db.Committee, error)
	GetCommitteeById(ctx context.Context, id uuid.UUID) (db.Committee, error)
}

type committeeService struct {
	queries db.Queries
}

func NewCommitteeService(queries db.Queries) CommitteeService {
	return &committeeService{queries: queries}
}

func (s *committeeService) GetAllCommittees(ctx context.Context) ([]db.Committee, error) {
	return s.queries.ListCommittees(ctx)
}

func (s *committeeService) CreateCommittee(ctx context.Context, committee db.CreateCommitteeParams) (db.Committee, error) {
	return s.queries.CreateCommittee(ctx, committee)
}

func (s *committeeService) GetCommitteeById(ctx context.Context, id uuid.UUID) (db.Committee, error) {
	return s.queries.GetCommittee(ctx, id)
}
