/*
Package types contains data structures and interfaces for the API server.
It contains a struct User and a struct LoginUserPayload that are used in the API server.
*/
package types

import (
	"net/http"
	"time"
)

type Middleware interface {
	RateLimit(limit int, window time.Duration) func(http.Handler) http.Handler
}

type Auth interface {
	HashPassword(password string) (string, error)
	ComparePassword(hash string, plain []byte) bool
	ConfirmPassword(password string, confirmPassword string) bool
	CreateJWT(secret []byte, userID int) (string, error)
}

// UserStore is an interface for the user store.
// It contains methods for getting and creating users.
type UserStore interface {
	GetUserByEmail(email string) (*User, error) // Get user by email
	GetUserByID(id string) (*User, error)       // Get user by ID
	CreateUser(User) error                      // Create user
}

// User is a struct that represents a user.
type User struct {
	ID        int    `json:"id"`         // User ID
	Name      string `json:"name"`       // User name
	Email     string `json:"email"`      // User email
	Password  string `json:"password"`   // User password
	CreatedAt string `json:"created_at"` // User creation date
	UpdatedAt string `json:"updated_at"` // User last updated
}

// LoginUserPayload is a struct that represents a login user payload.
type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email,min=2,max=320"` // User email
	Password string `json:"password" validate:"required,min=6,max=100"`    // User password
}

type RegisterUserPayload struct {
	Name            string `json:"name" validate:"required,min=2,max=100"`             // User name
	Email           string `json:"email" validate:"required,email,max=320"`            // User email
	Password        string `json:"password" validate:"required,min=6,max=100"`         // User password
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6,max=100"` // User confirm password
}
