package db

import (
	"testing"
)

func TestRedisFormatURL(t *testing.T) {
	rdb := RedisConfig{
		Host:     "redis",
		Port:     "6379",
		Password: "password",
		DB:       0,
	}

	expectedURL := "redis:6379"
	if url := rdb.FormatURL(); url != expectedURL {
		t.Errorf("Expected %s, but got %s", expectedURL, url)
	}
}
