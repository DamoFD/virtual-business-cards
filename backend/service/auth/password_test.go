package auth

import (
	"testing"
)

var a = NewAuthService()

func TestHashPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := a.HashPassword(password)
	if err != nil {
		t.Errorf("HashPassword() error = %v", err)
	}

	if hashedPassword == "" {
		t.Errorf("Expected hashed password, got empty string")
	}

	if hashedPassword == password {
		t.Errorf("Expected hashed password to be different from original password")
	}
}

func TestComparePassword(t *testing.T) {
	password := "password"
	hashedPassword, err := a.HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	if !a.ComparePassword(hashedPassword, []byte(password)) {
		t.Errorf("Expected password to match hashed password")
	}

	if a.ComparePassword(hashedPassword, []byte("wrong-password")) {
		t.Errorf("Expected password to not match hashed password")
	}
}

func TestConfirmPassword(t *testing.T) {
	password := "password"
	confirmPassword := "password"
	t.Run("passwords match", func(t *testing.T) {
		if !a.ConfirmPassword(password, confirmPassword) {
			t.Errorf("Expected passwords to match")
		}
	})

	t.Run("passwords don't match", func(t *testing.T) {
		if a.ConfirmPassword(password, "wrong-password") {
			t.Errorf("Expected passwords to not match")
		}
	})
}
