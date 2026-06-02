package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/mapper"
	"github.com/itsektionen/mimer/model"
	"github.com/jackc/pgx/v5/pgtype"
)

type TrusteeRepository interface {
	ListByGroupID(ctx context.Context, groupID uuid.UUID) ([]model.Trustee, error)
	ListAll(ctx context.Context) ([]model.Trustee, error)
	Create(ctx context.Context, positionID uuid.UUID, personID uuid.UUID, startDate time.Time, endDate time.Time) (*model.Trustee, error)
}

type trusteeRepository struct {
	q db.Querier
}

func NewTrusteeRepository(q db.Querier) TrusteeRepository {
	return &trusteeRepository{q: q}
}

func (r *trusteeRepository) ListByGroupID(ctx context.Context, groupID uuid.UUID) ([]model.Trustee, error) {
	trustees, err := r.q.ListGroupTrustees(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return mapper.ToGroupTrustees(trustees), nil
}

func (r *trusteeRepository) ListAll(ctx context.Context) ([]model.Trustee, error) {
	trustees, err := r.q.ListTrustees(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.ToListTrustees(trustees), nil
}

func (r *trusteeRepository) Create(ctx context.Context, positionID uuid.UUID, userID uuid.UUID, startDate time.Time, endDate time.Time) (*model.Trustee, error) {
	trustee, err := r.q.CreateTrustee(ctx, db.CreateTrusteeParams{
		PositionID: positionID,
		UserID:     userID,
		StartDate: pgtype.Date{
			Time:  startDate,
			Valid: true,
		},
		EndDate: pgtype.Date{
			Time:  endDate,
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	person, err := r.q.GetUser(ctx, trustee.UserID)
	if err != nil {
		return nil, err
	}

	position, err := r.q.GetPosition(ctx, trustee.PositionID)
	if err != nil {
		return nil, err
	}

	resp := mapper.TrusteeFromDB(trustee, person, position)

	return &resp, nil
}
