/*
Package api contains API server configurations and methods.
It contains a struct APIServer and a method Run() that starts the API server.

Example:

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
	    log.Fatal(err)
	}
*/
package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"

	"github.com/DamoFD/virtual-business/service/auth"
	"github.com/DamoFD/virtual-business/service/card"
	"github.com/DamoFD/virtual-business/service/middleware"
	"github.com/DamoFD/virtual-business/service/user"
)

// APIServer contains servers configurations.
type APIServer struct {
	addr string  // addr is the address of the server.
	db   *sql.DB // db is the database connection.
	rdb  *redis.Client
}

// NewAPIServer creates a new APIServer instance.
// It takes an address and a database connection as arguments.
// It returns an APIServer instance.
func NewAPIServer(addr string, db *sql.DB, rdb *redis.Client) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
		rdb:  rdb,
	}
}

// Run starts the API server.
// It returns an error if the server fails to start.
// It returns nil if the server starts successfully.
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	auth := auth.NewAuthService(s.rdb)
	middleware := middleware.NewMiddlewareService(s.rdb)

	// Register user routes
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore, auth)
	userHandler.RegisterRoutes(subrouter, middleware)

	// Register card routes
	cardStore := card.NewStore(s.db)
	cardHandler := card.NewHandler(cardStore, auth)
	cardHandler.RegisterRoutes(subrouter, middleware)

	log.Println("Listening on", s.addr)

    enhancedRouter := enableCORS(router)

	return http.ListenAndServe(s.addr, enhancedRouter)
}

func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Set CORS headers
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        w.Header().Set("Access-Control-Allow-Credentials", "true")

        // Check if the request is for CORS preflight
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}
