package middleware

import (
	"context"
	"net/http"
	"order-api/config"
	"order-api/pkg/jwt"
	"strings"
)

type key string

const (
	ContextUserPhone key = "userPhone"
)

func Auth(next http.Handler, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !strings.HasPrefix(authHeader, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		isValid, data := jwt.NewJWT(cfg).VerifyToken(token)
		if !isValid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), ContextUserPhone, data.UserPhone)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
