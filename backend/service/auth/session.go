package auth

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"

	"github.com/DamoFD/virtual-business/types"
)

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
