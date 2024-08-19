/*
Package types contains data structures and interfaces for the API server.
It contains a struct User and a struct LoginUserPayload that are used in the API server.
*/
package types

import (
	"context"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisClient is an interface for the Redis client.
type RedisClient interface {
	Incr(ctx context.Context, key string) *redis.IntCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
	Set(ctx context.Context, sessionID string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, sessionID string) *redis.StringCmd
	Del(ctx context.Context, sessionID ...string) *redis.IntCmd
}

// Middleware is an interface for the middleware.
type Middleware interface {
	RateLimit(limit int, window time.Duration) func(http.Handler) http.Handler
	Auth() func(http.Handler) http.Handler
}

// Auth is an interface for the authentication.
type Auth interface {
	HashPassword(password string) (string, error)
	ComparePassword(hash string, plain []byte) bool
	ConfirmPassword(password string, confirmPassword string) bool
	SetSession(ctx context.Context, u *User, expiration time.Duration) (string, error)
	GetSession(ctx context.Context, sessionID string) (*SessionData, error)
	DeleteSession(ctx context.Context, sessionID string) error
}

// UserStore is an interface for the user store.
// It contains methods for getting and creating users.
type UserStore interface {
	GetUserByEmail(email string) (*User, error) // Get user by email
	GetUserByID(id string) (*User, error)       // Get user by ID
	CreateUser(User) (User, error)              // Create user
}

type CardStore interface {
	GetCardBySlug(slug string) (*CardResponse, error)
	GetCardsByUserID(userID int) ([]*Card, error)
	CreateCard(ctx context.Context, card Card) (Card, error)
	CreateCardItemField(ctx context.Context, f CardItemField) error
	UpdateCard(ctx context.Context, card Card) (Card, error)
	DeleteCard(ctx context.Context, cardID int) error
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

type Card struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CardResponse struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Slug      string          `json:"slug"`
	UserID    int             `json:"user_id"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
	CardItems []FieldResponse `json:"fields"`
}

type FieldResponse struct {
	ID         int    `json:"id"`
	CardID     int    `json:"card_id"`
	CardItemID int    `json:"card_item_id"`
	Name       string `json:"name"`
	Value      string `json:"value"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type CardPayload struct {
	Name      string         `json:"name" validate:"required,min=2,max=100"`
	Slug      string         `json:"slug" validate:"required,min=2,max=100,slug"`
	CardItems []FieldPayload `json:"fields" validate:"required,dive"`
}

type FieldPayload struct {
	CardItemID int    `json:"card_item_id" validate:"required,gt=0"`
	Name       string `json:"name" validate:"required,min=2,max=100"`
	Value      string `json:"value" validate:"required,min=2,max=100"`
}

type CardItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CardItemField struct {
	ID         int    `json:"id"`
	CardID     int    `json:"card_id"`
	CardItemID int    `json:"card_item_id"`
	Name       string `json:"name"`
	Value      string `json:"value"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type SessionData struct {
	UserID    int    `json:"user_id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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
