package auth

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/DamoFD/virtual-business/types"
)

type MockRedisClient struct {
	data map[string]string
}

func (m *MockRedisClient) Incr(ctx context.Context, key string) *redis.IntCmd {
	val, ok := m.data[key]
	if !ok {
		m.data[key] = "1"
		return redis.NewIntResult(1, nil)
	}
	count := stringToInt(val) + 1
	m.data[key] = intToString(count)
	return redis.NewIntResult(int64(count), nil)
}

func (m *MockRedisClient) Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return redis.NewBoolResult(true, nil)
}

func (m *MockRedisClient) Set(ctx context.Context, sessionID string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return redis.NewStatusResult("OK", nil)
}

func stringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func intToString(i int) string {
	return strconv.Itoa(i)
}

func NewMockRedisClient() types.RedisClient {
	return &MockRedisClient{
		data: make(map[string]string),
	}
}

var mockRedis = NewMockRedisClient()

var a = NewAuthService(mockRedis)

func TestHashPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := a.HashPassword(password)
	if err != nil {
		t.Errorf("HashPassword() error = %v", err)
	}

	if hashedPassword == "" {
		t.Errorf("Expected hashed password, got empty string")
	}

	if hashedPassword == password {
		t.Errorf("Expected hashed password to be different from original password")
	}
}

func TestComparePassword(t *testing.T) {
	password := "password"
	hashedPassword, err := a.HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	if !a.ComparePassword(hashedPassword, []byte(password)) {
		t.Errorf("Expected password to match hashed password")
	}

	if a.ComparePassword(hashedPassword, []byte("wrong-password")) {
		t.Errorf("Expected password to not match hashed password")
	}
}

func TestConfirmPassword(t *testing.T) {
	password := "password"
	confirmPassword := "password"
	t.Run("passwords match", func(t *testing.T) {
		if !a.ConfirmPassword(password, confirmPassword) {
			t.Errorf("Expected passwords to match")
		}
	})

	t.Run("passwords don't match", func(t *testing.T) {
		if a.ConfirmPassword(password, "wrong-password") {
			t.Errorf("Expected passwords to not match")
		}
	})
}
