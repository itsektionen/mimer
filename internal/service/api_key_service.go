package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"github.com/itsektionen/mimer/internal/model"
	"github.com/itsektionen/mimer/internal/repository"
)

type ApiKeyService interface {
	CreateApiKey(ctx context.Context) (model.ApiKey, error)
	GetByValue(ctx context.Context, value string) (model.ApiKey, error)
}

type apiKeyService struct {
	repo repository.ApiKeyRepository
}

func NewApiKeyService(repo repository.ApiKeyRepository) ApiKeyService {
	return &apiKeyService{repo: repo}
}

func generateRandomString() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (s *apiKeyService) CreateApiKey(ctx context.Context) (model.ApiKey, error) {
	key, err := generateRandomString()
	if err != nil {
		return model.ApiKey{}, err
	}

	return s.repo.Create(ctx, key)
}

func (s *apiKeyService) GetByValue(ctx context.Context, value string) (model.ApiKey, error) {
	return s.repo.GetByValue(ctx, value)
}
