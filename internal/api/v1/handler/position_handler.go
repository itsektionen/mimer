package handler

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/model"
	"github.com/itsektionen/mimer/internal/service"
)

type PositionHandler struct {
	positionService service.PositionService
}

func NewPositionHandler(s service.PositionService) *PositionHandler {
	return &PositionHandler{positionService: s}
}

type CreatePositionRequest struct {
	Body struct {
		Name        string    `json:"name"`
		Email       string    `json:"email"`
		CommitteeID uuid.UUID `json:"committee_id"`
	}
}

type CreatePositionResponse struct {
	Body model.Position
}

func (h *PositionHandler) HandleCreatePosition(ctx context.Context, input *CreatePositionRequest) (*CreatePositionResponse, error) {
	resp := &CreatePositionResponse{}
	newPosition := db.CreatePositionParams{
		Name:        input.Body.Name,
		Email:       input.Body.Email,
		CommitteeID: input.Body.CommitteeID,
	}

	position, err := h.positionService.CreatePosition(ctx, newPosition)
	if err != nil {
		return nil, err
	}

	resp.Body = position
	return resp, nil
}

type ListPositionsResponse struct {
	Body []model.Position
}

func (h *PositionHandler) HandleListPositions(ctx context.Context, input *struct{}) (*ListPositionsResponse, error) {
	resp := &ListPositionsResponse{}
	positions, err := h.positionService.GetAllPositions(ctx)
	if err != nil {
		return nil, err
	}

	resp.Body = positions
	return resp, nil
}

type GetPositionByIdRequest struct {
	ID uuid.UUID `path:"id"`
}

type GetPositionByIdResponse struct {
	Body model.Position
}

func (h *PositionHandler) HandleGetPositionById(ctx context.Context, input *GetPositionByIdRequest) (*GetPositionByIdResponse, error) {
	resp := &GetPositionByIdResponse{}

	position, err := h.positionService.GetPositionById(ctx, input.ID)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	resp.Body = position
	return resp, nil
}

type AssignPositionRequest struct {
	PositionID uuid.UUID `path:"id"`
	Body       struct {
		PersonID uuid.UUID `json:"personId"`
	}
}

type AssignPositionResponse struct {
	Body model.Trustee
}

func (h *PositionHandler) HandleAssignPosition(ctx context.Context, input *AssignPositionRequest) (*AssignPositionResponse, error) {
	resp := &AssignPositionResponse{}

	trustee, err := h.positionService.AssignPosition(ctx, input.PositionID, input.Body.PersonID)
	if err != nil {
		return nil, err
	}

	resp.Body = *trustee
	return resp, nil
}
