package user

import (
	"database/sql"

	"github.com/DamoFD/virtual-business/types"
)

// Store is a struct that contains a database connection.
// It contains methods for getting and creating users.
type Store struct {
	db *sql.DB
}

// NewStore creates a new user store.
// It takes a database connection as a parameter.
// It returns a pointer to the Store struct.
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// GetUserByEmail pulls a user from the database by email.
// It takes an email string as a parameter.
// It returns a pointer to the User struct and an error.
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

// GetUserByID pulls a user from the database by ID.
// It takes an ID string as a parameter.
// It returns a pointer to the User struct and an error.
func (s *Store) GetUserByID(id string) (*types.User, error) {
	return nil, nil
}

// CreateUser creates a new user in the database.
// It takes a User struct as a parameter.
// It returns an error if the creation fails.
func (s *Store) CreateUser(u types.User) error {
	return nil
}
