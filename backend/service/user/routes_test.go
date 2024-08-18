package user

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"

	"github.com/DamoFD/virtual-business/types"
)

type mockAuth struct {
	HashPasswordFn    func(password string) (string, error)
	ComparePasswordFn func(hash string, plain []byte) bool
	ConfirmPasswordFn func(password string, confirmPassword string) bool
	SetSessionFn      func(ctx context.Context, u *types.User, expiration time.Duration) (string, error)
}

func (m *mockAuth) HashPassword(password string) (string, error) {
	return m.HashPasswordFn(password)
}

func (m *mockAuth) ComparePassword(hash string, plain []byte) bool {
	return m.ComparePasswordFn(hash, plain)
}

func (m *mockAuth) ConfirmPassword(password string, confirmPassword string) bool {
	return m.ConfirmPasswordFn(password, confirmPassword)
}

func (m *mockAuth) SetSession(ctx context.Context, u *types.User, expiration time.Duration) (string, error) {
	return m.SetSessionFn(ctx, u, expiration)
}

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

func (m *mockUserStore) CreateUser(u types.User) (types.User, error) {
	return types.User{}, m.createUserErr
}

func sendPayload(payload interface{}, route string, t *testing.T) *httptest.ResponseRecorder {
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
		SetSessionFn: func(ctx context.Context, u *types.User, expiration time.Duration) (string, error) {
			return "session_id", nil
		},
	}

	mockStore := &mockUserStore{
		user: nil,
	}

	handler := &Handler{store: mockStore, auth: mockAuth}

	marshalled, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", route, bytes.NewBuffer(marshalled))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()

	router.HandleFunc(route, handler.handleRegister).Methods("POST")
	router.ServeHTTP(rr, req)

	return rr
}
