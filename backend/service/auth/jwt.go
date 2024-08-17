package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/DamoFD/virtual-business/config"
)

// contextKey is a value for use with context.WithValue.
type contextKey string

// UserKey is the key for the user in the context.
const UserKey contextKey = "userID"

// CreateJWT creates a new JWT token with an expiration time.
// It takes a secret key and a user ID as parameters.
// It returns a string and an error.
func (a *AuthService) CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
