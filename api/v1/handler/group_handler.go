package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/model"
	"github.com/itsektionen/mimer/service"
)

type GroupHandler struct {
	groupService service.GroupService
}

func NewGroupHandler(s service.GroupService) *GroupHandler {
	return &GroupHandler{groupService: s}
}

type CreateGroupRequest struct {
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
	Inactive bool `query:"inactive" doc:"Include inactive groups"`
}

type ListGroupsResponse struct {
	Body []model.Group `json:"body"`
}

func (h *GroupHandler) HandleListGroups(ctx context.Context, input *ListCommiteesRequest) (*ListGroupsResponse, error) {
	resp := &ListGroupsResponse{}
	groups, err := h.groupService.List(ctx, input.Inactive)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	resp.Body = groups
	return resp, nil
}

type GetGroupByIdRequest struct {
	ID uuid.UUID `path:"id"`
}

type GetGroupByIdResponse struct {
	Body model.Group `json:"body"`
}

func (h *GroupHandler) HandleGetGroupById(ctx context.Context, input *GetGroupByIdRequest) (*GetGroupByIdResponse, error) {
	resp := &GetGroupByIdResponse{}
	group, err := h.groupService.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	resp.Body = group
	return resp, nil
}

type GetGroupTrusteesRequest struct {
	GroupID  uuid.UUID `path:"id"`
	Inactive bool      `query:"inactive" doc:"Include inactive trustees"`
}

type GetGroupTrusteesResponse struct {
	Body []model.Trustee `json:"body"`
}

func (h *GroupHandler) HandleListGroupTrustees(
	ctx context.Context,
	input *GetGroupTrusteesRequest,
) (*GetGroupTrusteesResponse, error) {
	resp := &GetGroupTrusteesResponse{}

	trustees, err := h.groupService.ListTrustees(ctx, input.GroupID, input.Inactive)
	if err != nil {
		return nil, err
	}

	fmt.Println("FAH")
	fmt.Println(trustees)

	resp.Body = trustees
	return resp, nil
}
