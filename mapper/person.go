package mapper

import (
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/model"
)

func PersonFromDB(p db.Person) model.Person {
	return model.Person{
		ID:        p.ID,
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
