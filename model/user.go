package model

import (
	"strings"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	FirstName string    `json:"firstName" example:"Ture"`
	LastName  string    `json:"lastName" example:"Teknolog"`
	ImageURL  *string   `json:"imageUrl"`
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) Initials() string {
	lastNames := strings.Split(u.LastName, " ")

	var initials strings.Builder
	initials.WriteString(string(u.FirstName[0]))
	for _, name := range lastNames {
		if len(name) > 0 {
			initials.WriteString(string(name[0]))
		}
	}
	return strings.ToUpper(initials.String())
}
