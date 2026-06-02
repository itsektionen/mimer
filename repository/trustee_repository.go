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
	ListByCommitteeID(ctx context.Context, committeeID uuid.UUID) ([]model.Trustee, error)
	ListAll(ctx context.Context) ([]model.Trustee, error)
	Create(ctx context.Context, positionID uuid.UUID, personID uuid.UUID, startDate time.Time, endDate time.Time) (*model.Trustee, error)
}

type trusteeRepository struct {
	q db.Querier
}

func NewTrusteeRepository(q db.Querier) TrusteeRepository {
	return &trusteeRepository{q: q}
}

func (r *trusteeRepository) ListByCommitteeID(ctx context.Context, committeeID uuid.UUID) ([]model.Trustee, error) {
	trustees, err := r.q.ListCommitteeTrustees(ctx, committeeID)
	if err != nil {
		return nil, err
	}
	return mapper.ToCommitteeTrustees(trustees), nil
}

func (r *trusteeRepository) ListAll(ctx context.Context) ([]model.Trustee, error) {
	trustees, err := r.q.ListTrustees(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.ToListTrustees(trustees), nil
}

func (r *trusteeRepository) Create(ctx context.Context, positionID uuid.UUID, personID uuid.UUID, startDate time.Time, endDate time.Time) (*model.Trustee, error) {
	trustee, err := r.q.CreateTrustee(ctx, db.CreateTrusteeParams{
		PositionID: positionID,
		PersonID:   personID,
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

	person, err := r.q.GetPerson(ctx, trustee.PersonID)
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
