package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Position struct {
	ID            uuid.UUID  `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	Name          string     `json:"name" example:"Bootloader"`
	GroupID       uuid.UUID  `json:"groupId" example:"00000000-0000-0000-0000-000000000000"`
	Email         string     `json:"email" example:"bootloader@kth.it"`
	EstablishedAt time.Time  `json:"establishedAt" example:"2021-01-01T00:00:00Z"`
	DissolvedAt   *time.Time `json:"dissolvedAt" example:"2021-01-01T00:00:00Z"`
}

func (p *Position) IsActive() bool {
	return p.DissolvedAt == nil || p.DissolvedAt.IsZero()
}

func (p *Position) Timespan() string {
	start := p.EstablishedAt.Format("2006")

	if p.IsActive() {
		return fmt.Sprintf("%s - Now", start)
	}

	end := p.DissolvedAt.Format("2006")
	return fmt.Sprintf("%s - %s", start, end)
}
