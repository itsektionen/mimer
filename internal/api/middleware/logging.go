package middleware

import (
	"fmt"
	"strconv"
	"time"

	"github.com/danielgtaylor/huma/v2"
)

// NOTE: Might seem weird to call next immediately, but otherwise the status code won't be known (since the request hasn't finished)
func LoggingMiddleware(ctx huma.Context, next func(ctx huma.Context)) {
	next(ctx)

	fmt.Printf(
		"%s [%s] %s %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		ctx.Method(),
		strconv.Itoa(ctx.Status()),
		ctx.URL().Path,
	)
}
