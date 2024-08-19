package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRateLimit(t *testing.T) {
	mockRedis := NewMockRedisClient()
	middlewareService := NewMiddlewareService(mockRedis)

	tests := []struct {
		name         string
		limit        int
		window       time.Duration
		requests     int
		expectedCode int
	}{
		{
			name:         "Allow request under limit",
			limit:        5,
			window:       time.Minute,
			requests:     4,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Block request over limit",
			limit:        2,
			window:       time.Minute,
			requests:     3,
			expectedCode: http.StatusTooManyRequests,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := middlewareService.RateLimit(tt.limit, tt.window)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			// run the tests
			for i := 0; i < tt.requests; i++ {
				req := httptest.NewRequest("GET", "/", nil)
				w := httptest.NewRecorder()
				handler.ServeHTTP(w, req)
				assert.Equal(t, tt.expectedCode, w.Code)
			}
		})
	}
}
