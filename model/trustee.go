package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Trustee struct {
	ID        uuid.UUID `json:"id"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	User      User      `json:"user"`
	Position  Position  `json:"position"`
}

func (t *Trustee) IsActive() bool {
	return t.StartDate.Before(time.Now()) && t.EndDate.After(time.Now())
}

func (t *Trustee) Timespan() string {
	start := t.StartDate.Format("2006")

	end := t.EndDate.Format("2006")
	return fmt.Sprintf("%s - %s", start, end)
}
