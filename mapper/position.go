package mapper

import (
	"time"

	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/model"
)

func ToPosition(p db.Position) model.Position {
	var dissolvedAt *time.Time
	if p.DissolvedAt.Valid {
		dissolvedAt = &p.DissolvedAt.Time
	}
	return model.Position{
		ID:            p.ID,
		Name:          p.Name,
		GroupID:       p.GroupID,
		Email:         p.Email,
		Description:   p.Description,
		EstablishedAt: p.EstablishedAt.Time,
		DissolvedAt:   dissolvedAt,
	}
}

func ToPositions(positions []db.Position) []model.Position {
	result := make([]model.Position, len(positions))
	for i, p := range positions {
		result[i] = ToPosition(p)
	}
	return result
}

func ToPositionWithTrustees(rows []db.ListGroupPositionsWithActiveTrusteesRow) []model.PositionWithTrustee {
	result := make([]model.PositionWithTrustee, len(rows))
	for i, row := range rows {
		var dissolvedAt *time.Time
		if row.PositionDissolvedAt.Valid {
			dissolvedAt = &row.PositionDissolvedAt.Time
		}

		pos := model.Position{
			ID:            row.PositionID,
			Name:          row.PositionName,
			GroupID:       row.GroupID,
			Email:         row.PositionEmail,
			Description:   row.PositionDescription,
			EstablishedAt: row.PositionEstablishedAt.Time,
			DissolvedAt:   dissolvedAt,
		}

		var trustee *model.Trustee
		if row.TrusteeID.Valid {
			var firstName, lastName string
			if row.FirstName != nil {
				firstName = *row.FirstName
			}
			if row.LastName != nil {
				lastName = *row.LastName
			}
			trustee = &model.Trustee{
				ID:        row.TrusteeID.Bytes,
				StartDate: row.StartDate.Time,
				EndDate:   row.EndDate.Time,
				User: model.User{
					ID:        row.UserID.Bytes,
					FirstName: firstName,
					LastName:  lastName,
					ImageURL:  row.UserImageUrl,
				},
				Position: pos,
			}
		}

		result[i] = model.PositionWithTrustee{
			Position: pos,
			Trustee:  trustee,
		}
	}
	return result
}
