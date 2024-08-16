package db

import (
	"testing"
)

func TestFormatURL(t *testing.T) {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	}

	expectedURL := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	if url := cfg.FormatURL(); url != expectedURL {
		t.Errorf("Expected %s, but got %s", expectedURL, url)
	}
}
