package handler

import (
	"net/http"
	"time"

	"github.com/itsektionen/mimer/internal/response"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	var health_response struct {
		Uptime time.Duration `json:"uptime"`
		Status string        `json:"status"`
	}

	health_response.Uptime = time.Duration(time.Since(startTime).Seconds())
	health_response.Status = "UP"

	response.RespondWithJSON(w, http.StatusOK, health_response)
}
