package mapper

import (
	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/model"
)

func PersonFromDB(p db.Person) model.Person {
	return model.Person{
		ID:        p.ID.String(),
		FirstName: p.FirstName,
		LastName:  p.LastName,
		ImageURL:  p.ImageUrl,
	}
}

func PeopleFromDB(people []db.Person) []model.Person {
	result := make([]model.Person, len(people))
	for i, p := range people {
		result[i] = PersonFromDB(p)
	}
	return result
}
