package model

import "github.com/google/uuid"

type Committee struct {
	ID          uuid.UUID `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	Slug        string    `json:"slug" example:"init"`
	Name        string    `json:"name" example:"init"`
	ShortName   string    `json:"shortName,omitempty" example:"init"`
	Description *string   `json:"description,omitempty" example:"init is always spelled in lowercase, regardless of context."`
	Color       string    `json:"color,omitempty" example:"#000000"`
	ImageUrl    *string   `json:"imageUrl,omitempty"`
	WebsiteUrl  *string   `json:"websiteUrl,omitempty" example:"https://init.kth.it"`
	Active      bool      `json:"active" example:"true"`
}
