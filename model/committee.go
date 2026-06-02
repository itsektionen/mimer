package model

import (
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
	ImageUrl      *string    `json:"imageUrl,omitempty"`
	WebsiteUrl    *string    `json:"websiteUrl,omitempty" example:"https://init.kth.it"`
	EstablishedAt time.Time  `json:"establishedAt" example:"2021-01-01"`
	DissolvedAt   *time.Time `json:"dissolvedAt" example:"2021-01-01"`
}

func (g *Group) IsActive() bool {
	return g.DissolvedAt == nil || g.DissolvedAt.IsZero()
}
