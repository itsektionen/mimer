package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/model"
)

type CommitteeRepository interface {
	GetAllCommittees() ([]model.Committee, error)
	CreateCommittee(position *model.Committee) (*model.Committee, error)
	GetCommitteeById(id uuid.UUID) (*model.Committee, error)
}

type committeeRepository struct {
	db *sql.DB
}

func NewCommitteeRepository(db *sql.DB) CommitteeRepository {
	return &committeeRepository{db: db}
}

func (r *committeeRepository) GetAllCommittees() ([]model.Committee, error) {
	query := `SELECT id, name, description, slug, color FROM committee`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all committees: %w", err)
	}
	defer rows.Close()

	var committees []model.Committee = []model.Committee{}
	for rows.Next() {
		var com model.Committee
		err := rows.Scan(&com.ID, &com.Name, &com.Description, &com.Slug, &com.Color)
		if err != nil {
			return nil, fmt.Errorf("failed to scan committee row: %w", err)
		}
		committees = append(committees, com)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return committees, nil
}

func (r *committeeRepository) GetCommitteeById(id uuid.UUID) (*model.Committee, error) {
	query := `SELECT id, name, description, slug, color FROM committee WHERE ID = $1`

	var com model.Committee
	err := r.db.QueryRow(query, id).Scan(&com.ID, &com.Name, &com.Description, &com.Slug, &com.Color)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get committee (%s): %w", id, err)
	}

	return &com, nil
}

func (r *committeeRepository) CreateCommittee(committee *model.Committee) (*model.Committee, error) {
	defaultColor := "#cc99ff"
	if committee.Color == nil {
		committee.Color = &defaultColor
	}
	query := `
	INSERT INTO committee (name, description, slug, color)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`

	err := r.db.QueryRow(query,
		committee.Name,
		committee.Description,
		committee.Slug,
		committee.Color,
	).Scan(&committee.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to create committee: %w", err)
	}
	return committee, nil

}
