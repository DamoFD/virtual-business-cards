package auth

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"

	"github.com/DamoFD/virtual-business/types"
)

// SetSession sets the session data for the given user.
// It takes a context.Context, a *types.User, and a time.Duration as parameters.
// It returns a string and an error.
func (a *AuthService) SetSession(ctx context.Context, u *types.User, expiration time.Duration) (string, error) {

	sessionData := types.SessionData{
		UserID:    u.ID,
		Email:     u.Email,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	data, err := json.Marshal(sessionData)
	if err != nil {
		return "", err
	}

	sessionID := uuid.New().String()

	err = a.rdb.Set(ctx, sessionID, data, expiration).Err()
	if err != nil {
		return "", err
	}
	return sessionID, nil
}

// GetSession returns the session data for the given session ID.
// It takes a context.Context and a string as parameters.
// It returns an error if the session ID is invalid or the data cannot be retrieved from Redis.
func (a *AuthService) GetSession(ctx context.Context, sessionID string) (*types.SessionData, error) {

	// Get session data from Redis
	data, err := a.rdb.Get(ctx, sessionID).Result()
	if err != nil {
		return nil, err
	}

	// Format session data
	var sessionData types.SessionData
	err = json.Unmarshal([]byte(data), &sessionData)
	if err != nil {
		return nil, err
	}

	return &sessionData, nil
}

// DeleteSession deletes the session data for the given session ID.
// It takes a context.Context and a string as parameters.
// It returns an error if the session ID is invalid or the data cannot be deleted from Redis.
func (a *AuthService) DeleteSession(ctx context.Context, sessionID string) error {
	return a.rdb.Del(ctx, sessionID).Err()
}
