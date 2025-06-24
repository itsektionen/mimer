package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/model"
)

type PersonRepository interface {
	GetAllPeople() ([]model.Person, error)
	CreatePerson(position *model.Person) (*model.Person, error)
	GetPersonById(id uuid.UUID) (*model.Person, error)
}

type personRepository struct {
	db *sql.DB
}

func NewPersonRepository(db *sql.DB) PersonRepository {
	return &personRepository{db: db}
}

func (r *personRepository) GetAllPeople() ([]model.Person, error) {
	query := `SELECT id, first_name, last_name, image_url FROM person`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all persons: %w", err)
	}
	defer rows.Close()

	var people []model.Person = []model.Person{}
	for rows.Next() {
		var person model.Person
		err := rows.Scan(
			&person.ID,
			&person.FirstName,
			&person.LastName,
			&person.ImageURL,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan person row: %w", err)
		}
		people = append(people, person)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return people, nil
}

func (r *personRepository) GetPersonById(id uuid.UUID) (*model.Person, error) {
	query := `SELECT id, first_name, last_name, image_url FROM person WHERE ID = $1`

	var person model.Person
	err := r.db.QueryRow(query, id).Scan(
		&person.ID,
		&person.FirstName,
		&person.LastName,
		&person.ImageURL,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get person (%s): %w", id, err)
	}

	return &person, nil
}

func (r *personRepository) CreatePerson(person *model.Person) (*model.Person, error) {
	query := `
	INSERT INTO person (first_name, last_name, image_url)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	err := r.db.QueryRow(query,
		person.FirstName,
		person.LastName,
		person.ImageURL,
	).Scan(&person.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to create person: %w", err)
	}
	return person, nil

}
