package middleware

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/DamoFD/virtual-business/types"
)

type MiddlewareService struct {
	rdb *redis.Client
}

func NewMiddlewareService(rdb *redis.Client) types.Middleware {
	return &MiddlewareService{
		rdb: rdb,
	}
}

func (m *MiddlewareService) RateLimit(limit int, window time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.Background()

			// Create a unique key for the client and route
			key := m.createKey(r)

			// Increment the request count for the client
			requestCount, err := m.rdb.Incr(ctx, key).Result()
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			// If this is the first request, set the expiration for the key
			if requestCount == 1 {
				m.rdb.Expire(ctx, key, window)
			}

			// Check if the request count exceeds the limit
			if int(requestCount) > limit {
				w.Header().Set("Retry-After", strconv.Itoa(int(window.Seconds())))
				http.Error(w, "Too many requests", http.StatusTooManyRequests)
				return
			}
		})
	}
}

func (m *MiddlewareService) createKey(r *http.Request) string {
	clientIP := r.RemoteAddr
	method := r.Method
	path := r.URL.Path
	key := "rate_limit:" + method + ":" + path + ":" + clientIP

	return key
}
