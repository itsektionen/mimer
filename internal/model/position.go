package model

type Position struct {
	ID          string `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	Name        string `json:"name" example:"Bootloader"`
	Active      bool   `json:"active" example:"true"`
	CommitteeID string `json:"committeeId" example:"00000000-0000-0000-0000-000000000000"`
	Email       string `json:"email" example:"bootloader@kth.it"`
}
