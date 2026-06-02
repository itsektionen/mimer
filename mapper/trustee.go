package mapper

import (
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/model"
)

func ToCommitteeTrustee(row db.ListCommitteeTrusteesRow) model.Trustee {
	return model.Trustee{
		ID:        row.TrusteeID,
		StartDate: row.StartDate.Time,
		EndDate:   row.EndDate.Time,
		Person: model.Person{
			ID:        row.PersonID,
			FirstName: row.FirstName,
			LastName:  row.LastName,
			ImageURL:  row.PersonImageUrl,
		},
		Position: model.Position{
			ID:          row.PositionID,
			Name:        row.PositionName,
			CommitteeID: row.CommitteeID,
			Email:       row.PositionEmail,
		},
	}
}

func ToCommitteeTrustees(rows []db.ListCommitteeTrusteesRow) []model.Trustee {
	result := make([]model.Trustee, len(rows))
	for i, row := range rows {
		result[i] = ToCommitteeTrustee(row)
	}
	return result
}

func TrusteeFromDB(t db.Trustee, p db.Person, pos db.Position) model.Trustee {
	return model.Trustee{
		ID:        t.ID,
		StartDate: t.StartDate.Time,
		EndDate:   t.EndDate.Time,
		Person:    PersonFromDB(p),
		Position:  PositionFromDB(pos),
	}
}

func toListTrustee(row db.ListTrusteesRow) model.Trustee {
	return model.Trustee{
		ID:        row.TrusteeID,
		StartDate: row.StartDate.Time,
		EndDate:   row.EndDate.Time,
		Person: model.Person{
			ID:        row.PersonID,
			FirstName: row.FirstName,
			LastName:  row.LastName,
			ImageURL:  nil,
		},
		Position: model.Position{
			ID:          row.PositionID,
			Name:        row.PositionName,
			Active:      true,            // TODO: Get actual active status
			CommitteeID: row.CommitteeID, // TODO: Replace with committee
			Email:       "",              // TODO: Get actual email
		},
	}
}

func ToListTrustees(rows []db.ListTrusteesRow) []model.Trustee {
	result := make([]model.Trustee, len(rows))
	for i, row := range rows {
		result[i] = toListTrustee(row)
	}

	return result
}
