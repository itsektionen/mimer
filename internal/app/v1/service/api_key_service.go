package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"github.com/itsektionen/mimer/internal/pkg/db"
)

type ApiKeyService interface {
	CreateApiKey(ctx context.Context) (db.ApiKey, error)
}

type apiKeyService struct {
	queries db.Queries
}

func NewApiKeyService(queries db.Queries) ApiKeyService {
	return &apiKeyService{queries: queries}
}

func generateRandomString() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (s *apiKeyService) CreateApiKey(ctx context.Context) (db.ApiKey, error) {
	key, err := generateRandomString()
	if err != nil {
		panic("fuck")
	}

	return s.queries.CreateApiKey(ctx, key)
}
