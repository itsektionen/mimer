package handler

import (
	"net/http"
	"time"

	"github.com/itsektionen/mimer/internal/pkg/util"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	var response struct {
		Uptime time.Duration `json:"uptime"`
		Status string        `json:"status"`
	}

	response.Uptime = time.Duration(time.Since(startTime).Seconds())
	response.Status = "UP"

	util.RespondWithJSON(w, http.StatusOK, response)
}
