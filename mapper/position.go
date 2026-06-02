package mapper

import (
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/model"
)

func ToPosition(p db.Position) model.Position {
	return model.Position{
		ID:          p.ID,
		Name:        p.Name,
		CommitteeID: p.CommitteeID,
		Email:       p.Email,
	}
}

func ToPositions(positions []db.Position) []model.Position {
	result := make([]model.Position, len(positions))
	for i, p := range positions {
		result[i] = ToPosition(p)
	}
	return result
}
