package model

import "time"

type Trustee struct {
	ID        string    `json:"id"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Person    Person    `json:"person"`
	Position  Position  `json:"position"`
}
