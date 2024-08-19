package auth

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/DamoFD/virtual-business/types"
)

func TestCreateSession(t *testing.T) {

	mockRedis := NewMockRedisClient()
	authService := NewAuthService(mockRedis)

	user := &types.User{
		ID:        1,
		Email:     "test@example.com",
		Name:      "Test User",
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	ctx := context.Background()
	expiration := time.Hour * 24

	sessionID, err := authService.SetSession(ctx, user, expiration)
	if err != nil {
		t.Errorf("SetSession() error = %v", err)
	}

	if sessionID == "" {
		t.Errorf("Expected session ID, got empty string")
	}

	storedData, exists := mockRedis.(*MockRedisClient).data[sessionID]
	if !exists {
		t.Errorf("Expected session data to be stored in Redis")
	}

	expectedData := types.SessionData{
		UserID:    user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	expectedJSON, err := json.Marshal(expectedData)
	if err != nil {
		t.Errorf("Failed to marshal expected data: %v", err)
	}

	assert.Equal(t, string(expectedJSON), storedData, "Stored session data does not match expected data")
}

func TestGetSession(t *testing.T) {

	mockRedis := NewMockRedisClient()
	authService := NewAuthService(mockRedis)

	user := &types.User{
		ID:        1,
		Email:     "test@example.com",
		Name:      "Test User",
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	ctx := context.Background()
	expiration := time.Hour * 24

	sessionID, err := authService.SetSession(ctx, user, expiration)
	if err != nil {
		t.Errorf("SetSession() error = %v", err)
	}

	sessionData, err := authService.GetSession(ctx, sessionID)
	if err != nil {
		t.Errorf("GetSession() error = %v", err)
	}

	if sessionData == nil {
		t.Errorf("Expected session data, got nil")
	}

	expectedData := &types.SessionData{
		UserID:    user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	assert.Equal(t, expectedData, sessionData, "Stored session data does not match expected data")
}
