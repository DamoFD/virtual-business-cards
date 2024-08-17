package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := HashPassword(password)
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
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	if !ComparePassword(hashedPassword, []byte(password)) {
		t.Errorf("Expected password to match hashed password")
	}

	if ComparePassword(hashedPassword, []byte("wrong-password")) {
		t.Errorf("Expected password to not match hashed password")
	}
}

func TestConfirmPassword(t *testing.T) {
	password := "password"
	confirmPassword := "password"
	t.Run("passwords match", func(t *testing.T) {
		if !ConfirmPassword(password, confirmPassword) {
			t.Errorf("Expected passwords to match")
		}
	})

	t.Run("passwords don't match", func(t *testing.T) {
		if ConfirmPassword(password, "wrong-password") {
			t.Errorf("Expected passwords to not match")
		}
	})
}
