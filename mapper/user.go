package mapper

import (
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/model"
)

func ToUser(p db.User) model.User {
	return model.User{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		ImageURL:  p.ImageUrl,
	}
}

func ToUsers(users []db.User) []model.User {
	result := make([]model.User, len(users))
	for i, p := range users {
		result[i] = ToUser(p)
	}
	return result
}
