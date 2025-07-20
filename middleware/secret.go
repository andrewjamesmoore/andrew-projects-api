package middleware

import (
	"context"
	"net/http"
)

type contextKey string

const SecretKey contextKey = "secret_key"

func WithSecretHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), SecretKey, r.Header.Get("x-secret-key"))
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetSecretKey(ctx context.Context) string {
	if val, ok := ctx.Value(SecretKey).(string); ok {
		return val
	}
	return ""
}
