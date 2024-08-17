package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"

	"github.com/DamoFD/virtual-business/types"
	"github.com/DamoFD/virtual-business/utils"
)

func TestHandleLogin(t *testing.T) {

	mockAuth := &mockAuth{
		HashPasswordFn: func(password string) (string, error) {
			return "$2a$10$abc123abc123abc123abcO/abc123abc123abc123abc123abc123abc", nil
		},
		ComparePasswordFn: func(hash string, plain []byte) bool {
			return true
		},
		ConfirmPasswordFn: func(password string, confirmPassword string) bool {
			return true
		},
		CreateJWTFn: func(secret []byte, userID int) (string, error) {
			return "token", nil
		},
	}

	validUser := &types.User{
		ID:        1,
		Name:      "Test User",
		Email:     "ZkX8x@example.com",
		Password:  "$2a$10$abc123abc123abc123abcO/abc123abc123abc123abc123abc123abc", // Mock bcrypt hashed password
		CreatedAt: "2022-01-01T00:00:00Z",
		UpdatedAt: "2022-01-01T00:00:00Z",
	}

	mockStore := &mockUserStore{
		user: validUser,
	}

	handler := &Handler{store: mockStore, auth: mockAuth}

	t.Run("invalid json", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/login", bytes.NewBufferString("{invalid json}"))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rr := httptest.NewRecorder()

		handler.handleLogin(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("unexpected status code: got %v want %v", status, http.StatusBadRequest)
		}
	})

	t.Run("validation error", func(t *testing.T) {
		invalidPayload := types.LoginUserPayload{Email: "invalid-email"}
		jsonPayload, _ := json.Marshal(invalidPayload)

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonPayload))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rr := httptest.NewRecorder()

		utils.Validate = validator.New()
		handler.handleLogin(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("unexpected status code: got %v want %v", status, http.StatusBadRequest)
		}
	})

	t.Run("user not found", func(t *testing.T) {
		mockStore.getUserErr = errors.New("user not found")

		validPayload := types.LoginUserPayload{Email: "ZkX8x@example.com", Password: "password"}
		jsonPayload, _ := json.Marshal(validPayload)

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonPayload))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rr := httptest.NewRecorder()

		handler.handleLogin(rr, req)

		if status := rr.Code; status != http.StatusUnauthorized {
			t.Errorf("unexpected status code: got %v want %v", status, http.StatusUnauthorized)
		}
	})

	t.Run("invalid password", func(t *testing.T) {
		mockStore.getUserErr = nil
		mockAuth.ComparePasswordFn = func(hash string, plain []byte) bool {
			return false
		}

		validPayload := types.LoginUserPayload{Email: "ZkX8x@example.com", Password: "password"}
		jsonPayload, _ := json.Marshal(validPayload)

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonPayload))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rr := httptest.NewRecorder()

		handler.handleLogin(rr, req)

		if status := rr.Code; status != http.StatusUnauthorized {
			t.Errorf("unexpected status code: got %v want %v", status, http.StatusUnauthorized)
		}
	})

	t.Run("successful login", func(t *testing.T) {
		mockAuth.ComparePasswordFn = func(hash string, plain []byte) bool {
			return true
		}

		mockAuth.CreateJWTFn = func(secret []byte, userID int) (string, error) {
			return "token", nil
		}

		validPayload := types.LoginUserPayload{Email: "ZkX8x@example.com", Password: "password"}
		jsonPayload, _ := json.Marshal(validPayload)

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonPayload))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rr := httptest.NewRecorder()

		handler.handleLogin(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("unexpected status code: got %v want %v", status, http.StatusOK)
		}

		expectedBody := `{"token":"token"}` + "\n"
		if rr.Body.String() != expectedBody {
			t.Errorf("unexpected body: got %v want %v", rr.Body.String(), expectedBody)
		}
	})
}
