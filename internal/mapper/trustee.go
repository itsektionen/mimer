package mapper

import (
	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/model"
)

func TrusteeFromDBRow(row db.ListCommitteeTrusteesRow) model.Trustee {
	return model.Trustee{
		ID:        row.TrusteeID.String(),
		StartDate: row.StartDate.Time,
		EndDate:   row.EndDate.Time,
		Person: model.Person{
			ID:        row.PersonID.String(),
			FirstName: row.FirstName,
			LastName:  row.LastName,
			ImageURL:  row.PersonImageUrl,
		},
		Position: model.Position{
			ID:          row.PositionID.String(),
			Name:        row.PositionName,
			Active:      row.PositionActive,
			CommitteeID: row.CommitteeID.String(),
			Email:       row.PositionEmail,
		},
	}
}

func TrusteesFromDBRows(rows []db.ListCommitteeTrusteesRow) []model.Trustee {
	result := make([]model.Trustee, len(rows))
	for i, row := range rows {
		result[i] = TrusteeFromDBRow(row)
	}
	return result
}

func TrusteeFromDB(t db.Trustee, p db.Person, pos db.Position) model.Trustee {
	return model.Trustee{
		ID:        t.ID.String(),
		StartDate: t.StartDate.Time,
		EndDate:   t.EndDate.Time,
		Person:    PersonFromDB(p),
		Position:  PositionFromDB(pos),
	}
}
