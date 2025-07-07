package service

import (
	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/model"
	"github.com/itsektionen/mimer/internal/repository"
)

type PositionService interface {
	GetAllPositions() ([]model.Position, error)
	CreatePosition(*model.Position) (*model.Position, error)
	GetPositionById(id uuid.UUID) (*model.Position, error)
}

func NewPositionService(positionRepo repository.PositionRepository) PositionService {
	return &positionService{positionRepo: positionRepo}
}

type positionService struct {
	positionRepo repository.PositionRepository
}

func (s *positionService) GetAllPositions() ([]model.Position, error) {
	return s.positionRepo.GetAllPositions()
}

func (s *positionService) CreatePosition(position *model.Position) (*model.Position, error) {
	return s.positionRepo.CreatePosition(position)
}

func (s *positionService) GetPositionById(id uuid.UUID) (*model.Position, error) {
	return s.positionRepo.GetPositionById(id)
}
