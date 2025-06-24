package model

import "time"

type Trustee struct {
	ID         string    `json:"id"`
	PersonID   string    `json:"personId"`
	PositionID string    `json:"positionId"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
}
