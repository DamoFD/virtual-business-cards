package db

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

type PostgresConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName string
    SSLMode string
}

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

func NewPostgreSQLStorage(cfg PostgresConfig) (*sql.DB, error) {
    url := cfg.FormatURL()
    db, err := sql.Open("postgres", url)
    if err != nil {
        log.Fatal(err)
    }

    return db, nil
}
