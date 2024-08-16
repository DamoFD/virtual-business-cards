package main

import (
    "log"
    "database/sql"

    "github.com/DamoFD/virtual-business/db"
    "github.com/DamoFD/virtual-business/config"
)

func main() {
    db, err := db.NewPostgreSQLStorage(db.PostgresConfig{
        Host: config.Envs.DBHost,
        Port: config.Envs.DBPort,
        User: config.Envs.DBUser,
        Password: config.Envs.DBPassword,
        SSLMode: config.Envs.DBSSLMode,
    })
    if err != nil {
        log.Fatal(err)
    }

    initStorage(db)
}

func initStorage(db *sql.DB) {
    err := db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Successfully connected!")
}
