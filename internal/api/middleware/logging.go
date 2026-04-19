package middleware

import (
	"fmt"
	"time"

	"github.com/danielgtaylor/huma/v2"
)

func LoggingMiddleware(ctx huma.Context, next func(ctx huma.Context)) {
	// TODO: Add status code to log, ctx.Status() is not working for some reason
	fmt.Printf(
		"%s [%s] %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		ctx.Method(),
		ctx.URL().Path,
	)
	next(ctx)
}
