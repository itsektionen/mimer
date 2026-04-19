package model

type Committee struct {
	ID          string  `json:"id"`
	Slug        string  `json:"slug"`
	Name        string  `json:"name"`
	ShortName   string  `json:"shortName,omitempty"`
	Description *string `json:"description,omitempty"`
	Color       string  `json:"color,omitempty"`
	ImageUrl    *string `json:"imageUrl,omitempty"`
	WebsiteUrl  *string `json:"websiteUrl,omitempty"`
}
