/*
Package config contains configurations for the API server.
It contains a struct Config and a method initConfig() that initializes the configurations.
*/
package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config contains configurations for the API server.
type Config struct {
	Port                   string // Port for the API server.
	DBHost                 string // Host for the database.
	DBPort                 string // Port for the database.
	DBUser                 string // User for the database.
	DBPassword             string // Password for the database.
	DBName                 string // Name for the database.
	DBSSLMode              string // SSL mode for the database.
	JWTExpirationInSeconds int64  // JWT expiration in seconds.
	JWTSecret              string // JWT secret key.
	RedisHost              string // Redis host
	RedisPort              string // Redis port
	RedisPassword          string // Redis password
	RedisDB                int64  // Redis database
}

// Envs contains configurations for the API server.
var Envs = initConfig()

// initConfig initializes the configurations.
// It loads the environment variables from the .env file.
// It returns a Config struct.
func initConfig() Config {
	godotenv.Load()

	return Config{
		Port:                   getEnv("PORT", "8080"),
		DBHost:                 getEnv("DB_HOST", "db"),
		DBPort:                 getEnv("DB_PORT", "5432"),
		DBUser:                 getEnv("DB_USER", "postgres"),
		DBPassword:             getEnv("DB_PASSWORD", "postgres"),
		DBName:                 getEnv("DB_NAME", "postgres"),
		DBSSLMode:              getEnv("DB_SSL_MODE", "disable"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "secret"),
		RedisHost:              getEnv("REDIS_HOST", "redis"),
		RedisPort:              getEnv("REDIS_PORT", "6379"),
		RedisPassword:          getEnv("REDIS_PASSWORD", ""),
		RedisDB:                getEnvAsInt("REDIS_DB", 0),
	}
}

// getEnv returns the value of the environment variable with the given key.
// If the variable is not set, it returns the fallback value.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// getEnvAsInt returns the value of the environment variable with the given key as an int64.
// If the variable is not set, it returns the fallback value.
func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
