package ctxs

import "context"

type Key string

var (
	pathKey  Key = "current_path"
	themeKey Key = "theme"
	tokenKey Key = "token"
)

func ThemeIntoContext(ctx context.Context, theme string) context.Context {
	return context.WithValue(ctx, themeKey, theme)
}

func ThemeFromContext(ctx context.Context) (string, bool) {
	theme, ok := ctx.Value(themeKey).(string)
	return theme, ok
}

func PathIntoContext(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, pathKey, path)
}

func PathFromContext(ctx context.Context) (string, bool) {
	path, ok := ctx.Value(pathKey).(string)
	return path, ok
}

func TokenIntoContext(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

func TokenFromContext(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(tokenKey).(string)
	return token, ok
}
