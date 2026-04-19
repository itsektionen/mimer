package middleware

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/itsektionen/mimer/internal/service"
)

type AuthMiddleware struct {
	api           huma.API
	apiKeyService service.ApiKeyService
}

func NewAuthMiddleware(api huma.API, apiKeyService service.ApiKeyService) *AuthMiddleware {
	return &AuthMiddleware{api: api, apiKeyService: apiKeyService}
}

func (m *AuthMiddleware) Handle(ctx huma.Context, next func(ctx huma.Context)) {

	authHeader := ctx.Header("Authorization")

	if ctx.Method() == http.MethodGet {
		next(ctx)
		return
	}

	apiKey, err := m.apiKeyService.GetByValue(ctx.Context(), authHeader)
	if err != nil {
		huma.WriteErr(m.api, ctx, http.StatusForbidden, "forbidden")
		return
	}

	if !apiKey.Active {
		huma.WriteErr(m.api, ctx, http.StatusForbidden, "forbidden")
		return
	}

	next(ctx)
}
