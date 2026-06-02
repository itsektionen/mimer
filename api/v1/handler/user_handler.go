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

type GetUserByIdRequest struct {
	ID uuid.UUID `path:"id"`
}

type GetUserByIdResponse struct {
	Body model.User
}

func (h *UserHandler) HandleGetUserById(ctx context.Context, input *GetUserByIdRequest) (*GetUserByIdResponse, error) {
	resp := &GetUserByIdResponse{}
	user, err := h.userService.GetByID(ctx, input.ID)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	resp.Body = user
	return resp, nil
}
