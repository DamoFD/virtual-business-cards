package middleware

import (
	"context"
	"strconv"
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

func (m *MockRedisClient) Get(ctx context.Context, sessionID string) *redis.StringCmd {
	return redis.NewStringResult(m.data[sessionID], nil)
}

func (m *MockRedisClient) Del(ctx context.Context, sessionID ...string) *redis.IntCmd {
	return redis.NewIntResult(1, nil)
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
