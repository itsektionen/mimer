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

	apiKey, err := m.apiKeyService.GetByValue(ctx.Context(), authHeader)
	if err != nil {
		_ = huma.WriteErr(m.api, ctx, http.StatusForbidden, "forbidden")
		return
	}

	if !apiKey.Active {
		_ = huma.WriteErr(m.api, ctx, http.StatusForbidden, "forbidden")
		return
	}

	next(ctx)
}
