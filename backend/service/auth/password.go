package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password with bcrypt.
// It takes a password string as a parameter.
// It returns a hashed password string and an error.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePasswords compares a hashed password with a plain password.
// It takes a hashed password string and a plain password string as parameters.
// It returns true if the passwords match, false otherwise.
var ComparePassword = func(hash string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), plain)
	return err == nil
}
