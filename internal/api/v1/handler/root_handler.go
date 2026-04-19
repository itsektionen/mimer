package handler

import (
	"net/http"

	"github.com/itsektionen/mimer/internal/response"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	response.RespondWithJSON(w, http.StatusOK, "Hello, world!")
}
