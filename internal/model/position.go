package model

import "github.com/google/uuid"

type Position struct {
	ID          uuid.UUID `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	Name        string    `json:"name" example:"Bootloader"`
	Active      bool      `json:"active" example:"true"`
	CommitteeID uuid.UUID `json:"committeeId" example:"00000000-0000-0000-0000-000000000000"`
	Email       string    `json:"email" example:"bootloader@kth.it"`
}
