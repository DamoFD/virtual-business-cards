package db

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

// RedisConfig contains configurations for the Redis database.
type RedisConfig struct {
	Host     string // REDIS_HOST
	Port     string // REDIS_PORT
	Password string // REDIS_PASSWORD
	DB       int64  // REDIS_DB
}

// FormatURL formats the URL for the Redis database.
// It takes a RedisConfig struct as an argument.
// It returns the formatted URL.
func (cfg RedisConfig) FormatURL() string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}

// NewRedisClient initializes the database.
// It takes a RedisConfig struct as an argument.
// It returns a *redis.Client and an error if the database fails to initialize.
func NewRedisClient(cfg RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.FormatURL(),
		Password: cfg.Password,
		DB:       int(cfg.DB),
	})

	return client, nil
}
