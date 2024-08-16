/*
Package db contains database configurations and methods.
It contains a struct PostgresConfig and a method NewPostgreSQLStorage() that initializes the database.
*/
package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// PostgresConfig contains configurations for the PostgreSQL database.
type PostgresConfig struct {
	Host     string // DB_HOST
	Port     string // DB_PORT
	User     string // DB_USER
	Password string // DB_PASSWORD
	DBName   string // DB_NAME
	SSLMode  string // DB_SSL_MODE
}

// FormatURL formats the URL for the PostgreSQL database.
// It takes a PostgresConfig struct as an argument.
// It returns the formatted URL.
func (cfg PostgresConfig) FormatURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
	)
}

// NewPostgreSQLStorage initializes the database.
// It takes a PostgresConfig struct as an argument.
// It returns a *sql.DB and an error if the database fails to initialize.
func NewPostgreSQLStorage(cfg PostgresConfig) (*sql.DB, error) {
	url := cfg.FormatURL()
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
