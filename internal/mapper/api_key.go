package mapper

import (
	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/model"
)

func ApiKeyFromDB(a db.ApiKey) model.ApiKey {
	return model.ApiKey{
		ID:     a.ID.String(),
		Value:  a.Value,
		Active: a.Active,
	}
}

func ApiKeysFromDB(apiKeys []db.ApiKey) []model.ApiKey {
	result := make([]model.ApiKey, len(apiKeys))
	for i, a := range apiKeys {
		result[i] = ApiKeyFromDB(a)
	}
	return result
}
