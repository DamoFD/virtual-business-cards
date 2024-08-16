/*
    Package main contains main function.
    It contains a function main that starts the API server and initializes the storage.
*/
package main

import (
    "log"
    "database/sql"

    "github.com/DamoFD/virtual-business/db"
    "github.com/DamoFD/virtual-business/config"
    "github.com/DamoFD/virtual-business/cmd/api"
)

// main starts the API server and initializes the storage.
// It returns an error if the server fails to start or the storage fails to initialize.
// It returns nil if the server and the storage start successfully.
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

    server := api.NewAPIServer(":8080", db)
    if err := server.Run(); err != nil {
        log.Fatal(err)
    }
}

// initStorage initializes the storage.
// It takes a database connection as an argument.
// It returns an error if the storage fails to initialize.
func initStorage(db *sql.DB) {
    err := db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Successfully connected!")
}
