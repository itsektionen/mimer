package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/mapper"
	"github.com/itsektionen/mimer/model"
)

type ApiKeyRepository interface {
	List(ctx context.Context) ([]model.ApiKey, error)
	GetByID(ctx context.Context, id uuid.UUID) (model.ApiKey, error)
	GetByValue(ctx context.Context, value string) (model.ApiKey, error)
	Create(ctx context.Context, value string) (model.ApiKey, error)
	Delete(ctx context.Context, id uuid.UUID) (model.ApiKey, error)
	Disable(ctx context.Context, id uuid.UUID) error
	Enable(ctx context.Context, id uuid.UUID) error
}

type apiKeyRepository struct {
	q db.Querier
}

func NewApiKeyRepository(q db.Querier) ApiKeyRepository {
	return &apiKeyRepository{q: q}
}

func (r *apiKeyRepository) List(ctx context.Context) ([]model.ApiKey, error) {
	apiKeys, err := r.q.ListApiKeys(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.ApiKeysFromDB(apiKeys), nil
}

func (r *apiKeyRepository) GetByID(ctx context.Context, id uuid.UUID) (model.ApiKey, error) {
	apiKey, err := r.q.GetApiKey(ctx, id)
	if err != nil {
		return model.ApiKey{}, err
	}
	return mapper.ApiKeyFromDB(apiKey), nil
}

func (r *apiKeyRepository) GetByValue(ctx context.Context, value string) (model.ApiKey, error) {
	apiKey, err := r.q.GetApiKeyByValue(ctx, value)
	if err != nil {
		return model.ApiKey{}, err
	}
	return mapper.ApiKeyFromDB(apiKey), nil
}

func (r *apiKeyRepository) Create(ctx context.Context, value string) (model.ApiKey, error) {
	apiKey, err := r.q.CreateApiKey(ctx, value)
	if err != nil {
		return model.ApiKey{}, err
	}
	return mapper.ApiKeyFromDB(apiKey), nil
}

func (r *apiKeyRepository) Delete(ctx context.Context, id uuid.UUID) (model.ApiKey, error) {
	apiKey, err := r.q.DeleteApiKey(ctx, id)
	if err != nil {
		return model.ApiKey{}, err
	}
	return mapper.ApiKeyFromDB(apiKey), nil
}

func (r *apiKeyRepository) Disable(ctx context.Context, id uuid.UUID) error {
	return r.q.DisableApiKey(ctx, id)
}

func (r *apiKeyRepository) Enable(ctx context.Context, id uuid.UUID) error {
	return r.q.EnableApiKey(ctx, id)
}
