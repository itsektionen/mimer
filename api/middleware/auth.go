package middleware

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/itsektionen/mimer/service"
)

func AuthMiddleware(api huma.API, apiKeyService service.ApiKeyService) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(ctx huma.Context)) {
		authHeader := ctx.Header("Authorization")

		apiKey, err := apiKeyService.GetByValue(ctx.Context(), authHeader)
		if err != nil {
			_ = huma.WriteErr(api, ctx, http.StatusForbidden, "forbidden")
			return
		}

		if !apiKey.Active {
			_ = huma.WriteErr(api, ctx, http.StatusForbidden, "forbidden")
			return
		}

		next(ctx)
	}
}
