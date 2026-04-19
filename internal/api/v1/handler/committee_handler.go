package handler

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/internal/db"
	"github.com/itsektionen/mimer/internal/model"
	"github.com/itsektionen/mimer/internal/service"
)

type CommitteeHandler struct {
	committeeService service.CommitteeService
}

func NewCommitteeHandler(s service.CommitteeService) *CommitteeHandler {
	return &CommitteeHandler{committeeService: s}
}

type CreateCommitteeRequest struct {
	Body struct {
		Name        string  `json:"name"`
		Slug        string  `json:"slug"`
		ShortName   string  `json:"shortName"`
		Description *string `json:"description,omitempty"`
		Color       string  `json:"color"`
		ImageUrl    *string `json:"imageUrl,omitempty"`
		WebsiteUrl  *string `json:"websiteUrl,omitempty"`
	}
}

type CreateCommitteeResponse struct {
	Body model.Committee
}

func (h *CommitteeHandler) HandleCreateCommittee(
	ctx context.Context,
	input *CreateCommitteeRequest,
) (*CreateCommitteeResponse, error) {
	resp := &CreateCommitteeResponse{}
	newCommittee := db.CreateCommitteeParams{
		Name:        input.Body.Name,
		Slug:        input.Body.Slug,
		ShortName:   input.Body.ShortName,
		Description: input.Body.Description,
		Color:       input.Body.Color,
		ImageUrl:    input.Body.ImageUrl,
		WebsiteUrl:  input.Body.WebsiteUrl,
	}

	committee, err := h.committeeService.CreateCommittee(ctx, newCommittee)
	if err != nil {
		return nil, err
	}

	resp.Body = committee
	return resp, nil
}

type ListCommiteesRequest struct {
	Inactive bool `query:"inactive" doc:"Include inactive committees"`
}

type ListCommitteesResponse struct {
	Body []model.Committee `json:"body"`
}

func (h *CommitteeHandler) HandleListCommittees(ctx context.Context, input *ListCommiteesRequest) (*ListCommitteesResponse, error) {
	resp := &ListCommitteesResponse{}
	committees, err := h.committeeService.GetAllCommittees(ctx, input.Inactive)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	resp.Body = committees
	return resp, nil
}

type GetCommitteeByIdRequest struct {
	ID uuid.UUID `path:"id"`
}

type GetCommitteeByIdResponse struct {
	Body model.Committee `json:"body"`
}

func (h *CommitteeHandler) HandleGetCommitteeById(ctx context.Context, input *GetCommitteeByIdRequest) (*GetCommitteeByIdResponse, error) {
	resp := &GetCommitteeByIdResponse{}
	committee, err := h.committeeService.GetCommitteeById(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	resp.Body = committee
	return resp, nil
}

type GetCommitteeTrusteesRequest struct {
	CommitteeID uuid.UUID `path:"id"`
	Inactive    bool      `query:"inactive" doc:"Include inactive trustees"`
}

type GetCommitteeTrusteesResponse struct {
	Body []model.Trustee `json:"body"`
}

func (h *CommitteeHandler) HandleListCommitteeTrustees(
	ctx context.Context,
	input *GetCommitteeTrusteesRequest,
) (*GetCommitteeTrusteesResponse, error) {
	resp := &GetCommitteeTrusteesResponse{}

	trustees, err := h.committeeService.GetCommitteeTrustees(ctx, input.CommitteeID, input.Inactive)
	if err != nil {
		return nil, err
	}
	resp.Body = trustees
	return resp, nil
}
