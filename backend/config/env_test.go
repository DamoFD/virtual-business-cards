package config

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	t.Run("environment variable is set", func(t *testing.T) {
		os.Setenv("TEST_ENV", "value")
		defer os.Unsetenv("TEST_ENV") // Clearn up after test

		got := getEnv("TEST_ENV", "default")
		want := "value"

		if got != want {
			t.Errorf("getEnv() = %q, want %q", got, want)
		}
	})

	t.Run("environment variable is not set", func(t *testing.T) {
		got := getEnv("TEST_ENV", "default")
		want := "default"

		if got != want {
			t.Errorf("getEnv() = %q, want %q", got, want)
		}
	})
}

func TestInitConfig(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("PORT", "8080")
	os.Setenv("DB_HOST", "db")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "postgres")
	os.Setenv("DB_SSL_MODE", "disable")

	// Clean up after the test
	defer os.Unsetenv("PORT")
	defer os.Unsetenv("DB_HOST")
	defer os.Unsetenv("DB_PORT")
	defer os.Unsetenv("DB_USER")
	defer os.Unsetenv("DB_PASSWORD")
	defer os.Unsetenv("DB_NAME")
	defer os.Unsetenv("DB_SSL_MODE")

	config := initConfig()

	// Check if the environment variables are set correctly
	if config.Port != "8080" {
		t.Errorf("initConfig().Port = %q, want %q", config.Port, "8080")
	}
	if config.DBHost != "db" {
		t.Errorf("initConfig().DBHost = %q, want %q", config.DBHost, "db")
	}
	if config.DBPort != "5432" {
		t.Errorf("initConfig().DBPort = %q, want %q", config.DBPort, "5432")
	}
	if config.DBUser != "postgres" {
		t.Errorf("initConfig().DBUser = %q, want %q", config.DBUser, "postgres")
	}
	if config.DBPassword != "postgres" {
		t.Errorf("initConfig().DBPassword = %q, want %q", config.DBPassword, "postgres")
	}
	if config.DBName != "postgres" {
		t.Errorf("initConfig().DBName = %q, want %q", config.DBName, "postgres")
	}
	if config.DBSSLMode != "disable" {
		t.Errorf("initConfig().DBSSLMode = %q, want %q", config.DBSSLMode, "disable")
	}
}

func TestInitConfigWithDefaults(t *testing.T) {
	// unset environment variables
	os.Unsetenv("PORT")
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("DB_USER")
	os.Unsetenv("DB_PASSWORD")
	os.Unsetenv("DB_NAME")
	os.Unsetenv("DB_SSL_MODE")

	config := initConfig()

	// Check if the config values match the default values
	if config.Port != "8080" {
		t.Errorf("initConfig().Port = %q, want %q", config.Port, "8080")
	}
	if config.DBHost != "db" {
		t.Errorf("initConfig().DBHost = %q, want %q", config.DBHost, "db")
	}
	if config.DBPort != "5432" {
		t.Errorf("initConfig().DBPort = %q, want %q", config.DBPort, "5432")
	}
	if config.DBUser != "postgres" {
		t.Errorf("initConfig().DBUser = %q, want %q", config.DBUser, "postgres")
	}
	if config.DBPassword != "postgres" {
		t.Errorf("initConfig().DBPassword = %q, want %q", config.DBPassword, "postgres")
	}
	if config.DBName != "postgres" {
		t.Errorf("initConfig().DBName = %q, want %q", config.DBName, "postgres")
	}
	if config.DBSSLMode != "disable" {
		t.Errorf("initConfig().DBSSLMode = %q, want %q", config.DBSSLMode, "disable")
	}
}
