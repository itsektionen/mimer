package model

import (
	"time"

	"github.com/google/uuid"
)

type Trustee struct {
	ID        uuid.UUID `json:"id"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Person    Person    `json:"person"`
	Position  Position  `json:"position"`
}

func (t *Trustee) IsActive() bool {
	return t.StartDate.Before(time.Now()) && t.EndDate.After(time.Now())
}
