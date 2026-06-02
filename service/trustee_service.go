package service

import (
	"context"

	"github.com/itsektionen/mimer/model"
	"github.com/itsektionen/mimer/repository"
)

type TrusteeService interface {
	ListActive(ctx context.Context) ([]model.Trustee, error)
}

type trusteeService struct {
	trusteeRepo repository.TrusteeRepository
}

func NewTrusteeService(
	trusteeRepo repository.TrusteeRepository,
) TrusteeService {
	return &trusteeService{
		trusteeRepo,
	}
}

func (s *trusteeService) ListActive(ctx context.Context) ([]model.Trustee, error) {
	trustees, err := s.trusteeRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	return trustees, nil
}
