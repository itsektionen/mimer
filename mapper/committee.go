package mapper

import (
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/model"
)

func CommitteeFromDB(c db.Committee) model.Committee {
	return model.Committee{
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

func CommitteesFromDB(committees []db.Committee) []model.Committee {
	result := make([]model.Committee, len(committees))
	for i, c := range committees {
		result[i] = CommitteeFromDB(c)
	}
	return result
}
