package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/mapper"
	"github.com/itsektionen/mimer/model"
)

type PersonRepository interface {
	List(ctx context.Context) ([]model.Person, error)
	GetByID(ctx context.Context, id uuid.UUID) (model.Person, error)
	Create(ctx context.Context, params db.CreatePersonParams) (model.Person, error)
	Update(ctx context.Context, params db.UpdatePersonParams) (model.Person, error)
	Delete(ctx context.Context, id uuid.UUID) (model.Person, error)
}

type personRepository struct {
	q db.Querier
}

func NewPersonRepository(q db.Querier) PersonRepository {
	return &personRepository{q: q}
}

func (r *personRepository) List(ctx context.Context) ([]model.Person, error) {
	people, err := r.q.ListPeople(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.PeopleFromDB(people), nil
}

func (r *personRepository) GetByID(ctx context.Context, id uuid.UUID) (model.Person, error) {
	person, err := r.q.GetPerson(ctx, id)
	if err != nil {
		return model.Person{}, err
	}
	return mapper.PersonFromDB(person), nil
}

func (r *personRepository) Create(ctx context.Context, params db.CreatePersonParams) (model.Person, error) {
	person, err := r.q.CreatePerson(ctx, params)
	if err != nil {
		return model.Person{}, err
	}
	return mapper.PersonFromDB(person), nil
}

func (r *personRepository) Update(ctx context.Context, params db.UpdatePersonParams) (model.Person, error) {
	person, err := r.q.UpdatePerson(ctx, params)
	if err != nil {
		return model.Person{}, err
	}
	return mapper.PersonFromDB(person), nil
}

func (r *personRepository) Delete(ctx context.Context, id uuid.UUID) (model.Person, error) {
	person, err := r.q.DeletePerson(ctx, id)
	if err != nil {
		return model.Person{}, err
	}
	return mapper.PersonFromDB(person), nil
}
