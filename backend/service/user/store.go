package user

import (
    "database/sql"

    "github.com/DamoFD/virtual-business/types"
)

type Store struct {
    db *sql.DB
}

func NewStore(db *sql.DB) *Store {
    return &Store{
        db: db,
    }
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
    return nil, nil
}
func (s *Store) GetUserByID(id string) (*types.User, error) {
    return nil, nil
}

func (s *Store) CreateUser(u types.User) error {
    return nil
}
