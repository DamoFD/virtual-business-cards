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

    "github.com/gorilla/mux"
)

// APIServer contains servers configurations.
type APIServer struct {
	addr string // addr is the address of the server.
	db   *sql.DB // db is the database connection.
}

// NewAPIServer creates a new APIServer instance.
// It takes an address and a database connection as arguments.
// It returns an APIServer instance.
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Run starts the API server.
// It returns an error if the server fails to start.
// It returns nil if the server starts successfully.
func (s *APIServer) Run() error {
    router := mux.NewRouter()
    // subrouter := router.PathPrefix("/api/v1").Subrouter()

    log.Println("Listening on", s.addr)

    return http.ListenAndServe(s.addr, router)
}
