package mapper

import (
	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/model"
)

func PositionFromDB(p db.Position) model.Position {
	return model.Position{
		ID:          p.ID,
		Name:        p.Name,
		Active:      p.Active,
		CommitteeID: p.CommitteeID,
		Email:       p.Email,
	}
}

func PositionsFromDB(positions []db.Position) []model.Position {
	result := make([]model.Position, len(positions))
	for i, p := range positions {
		result[i] = PositionFromDB(p)
	}
	return result
}
