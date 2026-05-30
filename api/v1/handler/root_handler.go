package handler

import (
	"net/http"

	"github.com/itsektionen/mimer/response"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	response.RespondWithJSON(w, http.StatusOK, "Hello, world!")
}
