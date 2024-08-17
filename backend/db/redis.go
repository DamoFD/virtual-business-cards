package db

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host     string // REDIS_HOST
	Port     string // REDIS_PORT
	Password string // REDIS_PASSWORD
	DB       int64  // REDIS_DB
}

func (cfg RedisConfig) FormatURL() string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}

func NewRedisClient(cfg RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.FormatURL(),
		Password: cfg.Password,
		DB:       int(cfg.DB),
	})

	return client, nil
}
