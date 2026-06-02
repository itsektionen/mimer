package mapper

import (
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/model"
)

func ToPosition(p db.Position) model.Position {
	return model.Position{
		ID:      p.ID,
		Name:    p.Name,
		GroupID: p.GroupID,
		Email:   p.Email,
	}
}

func ToPositions(positions []db.Position) []model.Position {
	result := make([]model.Position, len(positions))
	for i, p := range positions {
		result[i] = ToPosition(p)
	}
	return result
}
