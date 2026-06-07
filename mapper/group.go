package mapper

import (
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/model"
)

func ToGroup(c db.Group) model.Group {
	return model.Group{
		ID:            c.ID,
		Slug:          c.Slug,
		Name:          c.Name,
		ShortName:     c.ShortName,
		Description:   c.Description,
		ImageURL:      c.ImageUrl,
		WebsiteURL:    c.WebsiteUrl,
		Color:         c.Color,
		EstablishedAt: c.EstablishedAt.Time,
		DissolvedAt:   &c.DissolvedAt.Time,
	}
}

func ToGroups(groups []db.Group) []model.Group {
	result := make([]model.Group, len(groups))
	for i, c := range groups {
		result[i] = ToGroup(c)
	}
	return result
}
