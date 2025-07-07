package model

type Position struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Active      bool   `json:"active"`
	CommitteeID string `json:"committeeId"`
}
