/*
Package middleware contains middleware functions for the API server.
It contains a struct MiddlewareService and a method RateLimit() that limits the rate of requests.
It contains a method createKey() that creates a unique key for the client and route.
*/
package middleware

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/DamoFD/virtual-business/types"
)

// MiddlewareService is a struct that contains the redis client interface.
// It contains a method RateLimit() that limits the rate of requests.
// It contains a method createKey() that creates a unique key for the client and route.
type MiddlewareService struct {
	rdb types.RedisClient
}

// NewMiddlewareService initializes the middleware service.
// It takes a redis client interface as an argument.
// It returns a *MiddlewareService.
func NewMiddlewareService(rdb types.RedisClient) types.Middleware {
	return &MiddlewareService{
		rdb: rdb,
	}
}

// RateLimit limits the rate of requests.
// It takes a limit and a window as arguments.
// It returns a function that takes a http.Handler and returns a http.Handler.
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

			next.ServeHTTP(w, r)
		})
	}
}

// createKey creates a unique key for the client and route
// It takes a *http.Request as an argument.
// It returns a string.
func (m *MiddlewareService) createKey(r *http.Request) string {
	clientIP := r.RemoteAddr
	method := r.Method
	path := r.URL.Path
	key := "rate_limit:" + method + ":" + path + ":" + clientIP

	return key
}
