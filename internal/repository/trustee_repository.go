package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/mapper"
	"github.com/itsektionen/mimer/internal/model"
)

type TrusteeRepository interface {
	ListCommitteeTrustees(ctx context.Context, committeeID uuid.UUID) ([]model.Trustee, error)
}

type trusteeRepository struct {
	q db.Querier
}

func NewTrusteeRepository(q db.Querier) TrusteeRepository {
	return &trusteeRepository{q: q}
}

func (r *trusteeRepository) ListCommitteeTrustees(ctx context.Context, committeeID uuid.UUID) ([]model.Trustee, error) {
	trustees, err := r.q.ListCommitteeTrustees(ctx, committeeID)
	if err != nil {
		return nil, err
	}
	return mapper.TrusteesFromDBRows(trustees), nil
}
