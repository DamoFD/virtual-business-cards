package user

import (
	"database/sql"
	"fmt"

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
	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

// scanRowIntoUser scans a row from the database and returns a pointer to the User struct.
// It returns an error if the scan fails.
// It returns a pointer to the User struct and an error.
func scanRowIntoUser(row *sql.Rows) (*types.User, error) {
	u := new(types.User)
	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
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
