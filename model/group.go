package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Group struct {
	ID            uuid.UUID  `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	Slug          string     `json:"slug" example:"init"`
	Name          string     `json:"name" example:"init"`
	ShortName     string     `json:"shortName,omitempty" example:"init"`
	Description   *string    `json:"description,omitempty" example:"init is always spelled in lowercase, regardless of context."`
	Color         string     `json:"color,omitempty" example:"#000000"`
	ImageURL      *string    `json:"imageUrl,omitempty"`
	WebsiteURL    *string    `json:"websiteUrl,omitempty" example:"https://init.kth.it"`
	EstablishedAt time.Time  `json:"establishedAt" example:"2021-01-01"`
	DissolvedAt   *time.Time `json:"dissolvedAt,omitempty" example:"2021-01-01"`
}

func (g *Group) IsActive() bool {
	return g.DissolvedAt == nil || g.DissolvedAt.IsZero()
}

func (g *Group) Timespan() string {
	start := g.EstablishedAt.Format("2006")

	if g.IsActive() {
		return fmt.Sprintf("%s - Now", start)
	}

	end := g.DissolvedAt.Format("2006")
	return fmt.Sprintf("%s - %s", start, end)
}

type GroupWithPositions struct {
	Group     Group
	Positions []PositionWithTrustee
}
