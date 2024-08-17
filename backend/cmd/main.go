/*
Package main contains main function.
It contains a function main that starts the API server and initializes the storage.
*/
package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/go-redis/redis/v8"

	"github.com/DamoFD/virtual-business/cmd/api"
	"github.com/DamoFD/virtual-business/config"
	"github.com/DamoFD/virtual-business/db"
)

// main starts the API server and initializes the storage.
// It returns an error if the server fails to start or the storage fails to initialize.
// It returns nil if the server and the storage start successfully.
func main() {
	pg, err := db.NewPostgreSQLStorage(db.PostgresConfig{
		Host:     config.Envs.DBHost,
		Port:     config.Envs.DBPort,
		User:     config.Envs.DBUser,
		Password: config.Envs.DBPassword,
		SSLMode:  config.Envs.DBSSLMode,
	})
	if err != nil {
		log.Fatal(err)
	}
	initStorage(pg)

	rdb, err := db.NewRedisClient(db.RedisConfig{
		Host:     config.Envs.RedisHost,
		Port:     config.Envs.RedisPort,
		Password: config.Envs.RedisPassword,
		DB:       config.Envs.RedisDB,
	})
	if err != nil {
		log.Fatal(err)
	}
	initRedisStorage(rdb)

	server := api.NewAPIServer(":8080", pg, rdb)
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

	log.Println("Successfully connected to db!")
}

func initRedisStorage(client *redis.Client) {
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Successfully connected to redis!")
}
