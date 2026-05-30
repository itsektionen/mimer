package router

import (
	"net/http"
)

func SetupRootRouter(
	v1APIRouter http.Handler,
) http.Handler {

	return v1APIRouter
}
