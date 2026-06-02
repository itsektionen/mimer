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

func (s *trusteeService) ListAll(ctx context.Context) ([]model.Trustee, error) {
	trustees, err := s.trusteeRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	return trustees, nil
}

func (s *trusteeService) ListActive(ctx context.Context) ([]model.Trustee, error) {
	activeTrustees := make([]model.Trustee, 0)
	allTrustees, err := s.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, trustee := range allTrustees {
		if trustee.IsActive() {
			activeTrustees = append(activeTrustees, trustee)
		}
	}

	return activeTrustees, nil
}
