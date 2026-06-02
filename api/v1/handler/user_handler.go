package handler

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/model"
	"github.com/itsektionen/mimer/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{userService: s}
}

type ListUsersResponse struct {
	Body []model.User
}

func (h *UserHandler) HandleListUsers(ctx context.Context, input *struct{}) (*ListUsersResponse, error) {
	resp := &ListUsersResponse{}
	users, err := h.userService.List(ctx)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	resp.Body = users
	return resp, nil
}

type GetUserByIDRequest struct {
	ID uuid.UUID `path:"id"`
}

type GetUserByIDResponse struct {
	Body model.User
}

func (h *UserHandler) HandleGetUserByID(ctx context.Context, input *GetUserByIDRequest) (*GetUserByIDResponse, error) {
	resp := &GetUserByIDResponse{}
	user, err := h.userService.GetByID(ctx, input.ID)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	resp.Body = user
	return resp, nil
}
