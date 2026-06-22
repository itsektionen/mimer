package ctxs

import "context"

type Key string

var (
	PathKey  Key = "current_path"
	ThemeKey Key = "theme"
)

func ThemeFromContext(ctx context.Context) (string, bool) {
	theme, ok := ctx.Value(ThemeKey).(string)
	return theme, ok
}

func PathFromContext(ctx context.Context) (string, bool) {
	path, ok := ctx.Value(PathKey).(string)
	return path, ok
}
