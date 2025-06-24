package handler

import (
	"net/http"

	"github.com/itsektionen/mimer/internal/pkg/util"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusOK, "Hello, world!")
}
