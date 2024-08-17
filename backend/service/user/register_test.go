package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/DamoFD/virtual-business/types"
)

func TestHandleRegister(t *testing.T) {

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
		req, err := http.NewRequest("POST", "/register", bytes.NewBufferString("{invalid json}"))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rr := httptest.NewRecorder()

		handler.handleRegister(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})

	t.Run("validation error", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/register", bytes.NewBufferString("{\"name\":\"\",\"email\":\"\",\"password\":\"\"}"))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rr := httptest.NewRecorder()

		handler.handleRegister(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})

	t.Run("Passwords do not match", func(t *testing.T) {
		mockAuth.ConfirmPasswordFn = func(password string, confirmPassword string) bool {
			return false
		}

		req, err := http.NewRequest("POST", "/register", bytes.NewBufferString("{\"name\":\"\",\"email\":\"\",\"password\":\"password\",\"confirm_password\":\"not-matching\"}"))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rr := httptest.NewRecorder()

		handler.handleRegister(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})

	t.Run("User already exists", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/register", bytes.NewBufferString("{\"name\":\"Test User\",\"email\":\"ZkX8x@example.com\",\"password\":\"password\",\"confirm_password\":\"password\"}"))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rr := httptest.NewRecorder()

		handler.handleRegister(rr, req)

		if status := rr.Code; status != http.StatusConflict {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusConflict)
		}
	})

	t.Run("Success", func(t *testing.T) {

		mockAuth.ConfirmPasswordFn = func(password string, confirmPassword string) bool {
			return true
		}
		mockStore := &mockUserStore{
			user: nil,
		}
		handler := &Handler{store: mockStore, auth: mockAuth}

		payload := types.RegisterUserPayload{
			Name:            "New User",
			Email:           "test@example.com",
			Password:        "password",
			ConfirmPassword: "password",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister).Methods("POST")
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusCreated)
		}
	})
}
