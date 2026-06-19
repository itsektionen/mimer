package mapper

import (
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/model"
)

func toGroupTrustee(row db.ListGroupTrusteesRow) model.Trustee {
	return model.Trustee{
		ID:        row.TrusteeID,
		StartDate: row.StartDate.Time,
		EndDate:   row.EndDate.Time,
		User: model.User{
			ID:        row.UserID,
			FirstName: row.FirstName,
			LastName:  row.LastName,
			ImageURL:  row.UserImageUrl,
		},
		Position: model.Position{
			ID:      row.PositionID,
			Name:    row.PositionName,
			GroupID: row.GroupID,
			Email:   row.PositionEmail,
		},
	}
}

func ToGroupTrustees(rows []db.ListGroupTrusteesRow) []model.Trustee {
	result := make([]model.Trustee, len(rows))
	for i, row := range rows {
		result[i] = toGroupTrustee(row)
	}
	return result
}

func TrusteeFromDB(t db.Trustee, p db.User, pos db.Position) model.Trustee {
	return model.Trustee{
		ID:        t.ID,
		StartDate: t.StartDate.Time,
		EndDate:   t.EndDate.Time,
		User:      ToUser(p),
		Position:  ToPosition(pos),
	}
}

func toListTrustee(row db.ListTrusteesRow) model.Trustee {
	return model.Trustee{
		ID:        row.TrusteeID,
		StartDate: row.StartDate.Time,
		EndDate:   row.EndDate.Time,
		User: model.User{
			ID:        row.UserID,
			FirstName: row.FirstName,
			LastName:  row.LastName,
			ImageURL:  nil,
		},
		Position: model.Position{
			ID:      row.PositionID,
			Name:    row.PositionName,
			GroupID: row.GroupID, // TODO: Replace with group
			Email:   "",          // TODO: Get actual email
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

func ToLIstTrusteesByPosition(rows []db.ListTrusteesByPositionRow) []model.Trustee {
	result := make([]model.Trustee, len(rows))
	for i, row := range rows {
		result[i] = model.Trustee{
			ID:        row.TrusteeID,
			StartDate: row.StartDate.Time,
			EndDate:   row.EndDate.Time,
			User: model.User{
				ID:        row.UserID,
				FirstName: row.FirstName,
				LastName:  row.LastName,
				ImageURL:  nil,
			},
			Position: model.Position{
				ID:      row.PositionID,
				Name:    row.PositionName,
				GroupID: row.GroupID, // TODO: Replace with group
				Email:   "",          // TODO: Get actual email
			},
		}
	}

	return result
}
