package handler

import (
	"context"
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

type GetHealthResponse struct {
	Body struct {
		Uptime time.Duration `json:"uptime" example:"1776637039"`
		Status string        `json:"status" example:"UP"`
	}
}

func GetHealth(ctx context.Context, input *struct{}) (*GetHealthResponse, error) {
	resp := &GetHealthResponse{}

	resp.Body.Uptime = time.Duration(time.Since(startTime).Seconds())
	resp.Body.Status = "UP"

	return resp, nil
}
