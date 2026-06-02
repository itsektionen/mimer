package mapper

import (
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/model"
)

func ToGroup(c db.Committee) model.Group {
	return model.Group{
		ID:          c.ID,
		Slug:        c.Slug,
		Name:        c.Name,
		ShortName:   c.ShortName,
		Description: c.Description,
		ImageUrl:    c.ImageUrl,
		WebsiteUrl:  c.WebsiteUrl,
		Color:       c.Color,
		Active:      c.Active,
	}
}

func ToGroups(committees []db.Committee) []model.Group {
	result := make([]model.Group, len(committees))
	for i, c := range committees {
		result[i] = ToGroup(c)
	}
	return result
}
