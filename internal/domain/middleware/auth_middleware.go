package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/service"
)

type contextKey string

const (
	userIDKey contextKey = "userID"
	rolesKey  contextKey = "roles"
)

func AuthMiddleware(tokenService *service.TokenService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			header := r.Header.Get("Authorization")
			if header == "" {
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(header, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid authorization format", http.StatusUnauthorized)
				return
			}

			tokenStr := parts[1]

			claims, err := tokenService.ValidateAccessToken(tokenStr)
			if err != nil {
				http.Error(w, "invalid or expired token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userIDKey, claims.UserID)
			ctx = context.WithValue(ctx, rolesKey, claims.Roles)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
