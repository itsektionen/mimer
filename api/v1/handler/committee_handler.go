package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/model"
	"github.com/itsektionen/mimer/service"
)

type CommitteeHandler struct {
	committeeService service.GroupService
}

func NewCommitteeHandler(s service.GroupService) *CommitteeHandler {
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

type ListCommiteesRequest struct {
	Inactive bool `query:"inactive" doc:"Include inactive committees"`
}

type ListCommitteesResponse struct {
	Body []model.Group `json:"body"`
}

func (h *CommitteeHandler) HandleListCommittees(ctx context.Context, input *ListCommiteesRequest) (*ListCommitteesResponse, error) {
	resp := &ListCommitteesResponse{}
	committees, err := h.committeeService.List(ctx, input.Inactive)
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
	Body model.Group `json:"body"`
}

func (h *CommitteeHandler) HandleGetCommitteeById(ctx context.Context, input *GetCommitteeByIdRequest) (*GetCommitteeByIdResponse, error) {
	resp := &GetCommitteeByIdResponse{}
	committee, err := h.committeeService.GetByID(ctx, input.ID)
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

	trustees, err := h.committeeService.ListTrustees(ctx, input.CommitteeID, input.Inactive)
	if err != nil {
		return nil, err
	}

	fmt.Println("FAH")
	fmt.Println(trustees)

	resp.Body = trustees
	return resp, nil
}
