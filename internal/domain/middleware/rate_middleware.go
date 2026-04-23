package middleware

import (
	"net/http"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/service"
	"github.com/YagoSchramm/myecommerce-api/internal/infrastructure/datastore"
)

type RateLimitMiddleware struct {
	limiter datastore.Limiter
}

func NewRateLimitMiddleware(l datastore.Limiter) *RateLimitMiddleware {
	return &RateLimitMiddleware{limiter: l}
}

func (m *RateLimitMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var key string

		userID, ok := r.Context().Value(userIDKey).(string)
		if ok {
			key = "user:" + userID
		} else {
			ip := service.GetIp(r)
			key = "ip:" + ip
		}

		allowed, err := m.limiter.Allow(r.Context(), key)
		if err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		if !allowed {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
