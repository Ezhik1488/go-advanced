package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println(token)
		next.ServeHTTP(w, r)
	})
}
