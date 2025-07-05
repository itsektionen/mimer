package repository

import (
	"database/sql"
	"fmt"

	"github.com/itsektionen/mimer/internal/model"
)

type ApiKeyRepository interface {
	CreateApiKey(apiKey *model.ApiKey) (*model.ApiKey, error)
	GetApiKeyByValue(value string) (*model.ApiKey, error)
}

type apiKeyRepository struct {
	db *sql.DB
}

func NewApiKeyRepository(db *sql.DB) ApiKeyRepository {
	return &apiKeyRepository{db: db}
}

func (r *apiKeyRepository) CreateApiKey(apiKey *model.ApiKey) (*model.ApiKey, error) {
	query := `INSERT INTO api_key (value) VALUES $1 RETURNING value`

	err := r.db.QueryRow(query,
		apiKey.Value,
	).Scan(&apiKey.Value)

	if err != nil {
		return nil, fmt.Errorf("failed to create api key: %w", err)
	}

	return apiKey, nil
}

func (r *apiKeyRepository) GetApiKeyByValue(value string) (*model.ApiKey, error) {
	query := `SELECT * FROM api_key WHERE value = $1`

	var key model.ApiKey
	err := r.db.QueryRow(query,
		value,
	).Scan(&key.ID, &key.Value)

	if err != nil {
		return nil, fmt.Errorf("failed to get api key: %w", err)
	}

	return &key, nil
}
