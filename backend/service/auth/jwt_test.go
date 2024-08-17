package auth

import (
	"testing"
)

func TestCreateJWT(t *testing.T) {
	secret := []byte("secret")
	userID := 1
	token, err := a.CreateJWT(secret, userID)
	if err != nil {
		t.Errorf("CreateJWT() error = %v", err)
	}

	if token == "" {
		t.Errorf("Expected token, got empty string")
	}
}
