package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/DamoFD/virtual-business/service/auth"
	"github.com/DamoFD/virtual-business/types"
	"github.com/DamoFD/virtual-business/utils"
)

type mockUserStore struct {
	user          *types.User
	getUserErr    error
	createUserErr error
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	if m.user == nil {
		return nil, errors.New("user not found")
	}

	return m.user, m.getUserErr
}

func (m *mockUserStore) GetUserByID(id string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(u types.User) error {
	return m.createUserErr
}

func TestHandleRegister(t *testing.T) {

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

	handler := &Handler{store: mockStore}

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
		mockStore := &mockUserStore{
			user: nil,
		}
		handler := &Handler{store: mockStore}

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

func TestHandleLogin(t *testing.T) {
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

	handler := &Handler{store: mockStore}

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
		auth.ComparePassword = func(hash string, plain []byte) bool {
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
		auth.ComparePassword = func(hash string, plain []byte) bool {
			return true
		}

		auth.CreateJWT = func(secret []byte, userID int) (string, error) {
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
