package middleware

import (
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/zap"
)

// NOTE: Might seem weird to call next immediately, but otherwise the status code won't be known (since the request hasn't finished)
func LoggingMiddleware(logger *zap.Logger) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		start := time.Now()

		next(ctx)

		duration := time.Since(start)
		status := ctx.Status()
		if status == 0 {
			status = 200
		}

		fields := []zap.Field{
			zap.String("method", ctx.Method()),
			zap.String("path", ctx.URL().Path),
			zap.Int("status", status),
			zap.Duration("duration", duration),
		}

		// Choose level based on status
		if status >= http.StatusInternalServerError {
			logger.Error("request error", fields...)
		} else if status >= http.StatusBadRequest {
			logger.Warn("client error", fields...)
		} else {
			logger.Info("request handled", fields...)
		}
	}
}
