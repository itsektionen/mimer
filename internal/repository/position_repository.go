package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/model"
)

type PositionRepository interface {
	GetActivePositions() ([]model.Position, error)
	GetAllPositions() ([]model.Position, error)
	CreatePosition(position *model.Position) (*model.Position, error)
	GetPositionById(id uuid.UUID) (*model.Position, error)
}

type positionRepository struct {
	db *sql.DB
}

func NewPositionRepository(db *sql.DB) PositionRepository {
	return &positionRepository{db: db}
}

func (r *positionRepository) GetActivePositions() ([]model.Position, error) {
	query := "SELECT id, name, active, committee_id FROM position WHERE active = true"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all positions %w", err)
	}
	defer rows.Close()

	var positions []model.Position = []model.Position{}
	for rows.Next() {
		var position model.Position
		err := rows.Scan(
			&position.ID,
			&position.Name,
			&position.Active,
			&position.CommitteeID,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan position row: %w", err)
		}
		positions = append(positions, position)
	}
	return positions, nil
}

func (r *positionRepository) GetAllPositions() ([]model.Position, error) {
	query := "SELECT id, name, active, committee_id FROM position"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all positions %w", err)
	}
	defer rows.Close()

	var positions []model.Position = []model.Position{}
	for rows.Next() {
		var position model.Position
		err := rows.Scan(
			&position.ID,
			&position.Name,
			&position.Active,
			&position.CommitteeID,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan position row: %w", err)
		}
		positions = append(positions, position)
	}
	return positions, nil
}

func (r *positionRepository) CreatePosition(position *model.Position) (*model.Position, error) {
	query := `
	INSERT INTO position (name, active, committee_id)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	err := r.db.QueryRow(query,
		position.Name,
		position.Active,
		position.CommitteeID,
	).Scan(&position.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to create position: %w", err)
	}
	return position, nil
}

func (r *positionRepository) GetPositionById(id uuid.UUID) (*model.Position, error) {
	query := "SELECT id, name, active, committee_id FROM position WHERE id = $1"

	var position model.Position
	err := r.db.QueryRow(query, id).Scan(
		&position.ID,
		&position.Name,
		&position.Active,
		&position.CommitteeID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get position (%s): %w", id, err)
	}

	return &position, nil
}
