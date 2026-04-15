package middleware

import (
	"net/http"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/service"
)

func RequireRoles(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userRoles, ok := r.Context().Value(rolesKey).([]string)
			if !ok {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}

			if !service.HasAnyRole(userRoles, allowedRoles) {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
