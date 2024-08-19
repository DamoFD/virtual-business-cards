package user

import (
	"net/http"
	"testing"

	"github.com/DamoFD/virtual-business/types"
)

func TestValidateLoginPayload(t *testing.T) {

	testCases := []struct {
		name       string
		payload    types.LoginUserPayload
		wantStatus int
	}{
		{
			name: "Should return error if email is empty",
			payload: types.LoginUserPayload{
				Email:    "",
				Password: "password",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return error if password is empty",
			payload: types.LoginUserPayload{
				Email:    "test@example.com",
				Password: "",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return error if email is invalid",
			payload: types.LoginUserPayload{
				Email:    "test",
				Password: "password",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return error if email is greater than 320 characters",
			payload: types.LoginUserPayload{
				Email:    "fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff@gmail.com",
				Password: "password",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return error if password is less than 6 characters",
			payload: types.LoginUserPayload{
				Email:    "test@example.com",
				Password: "pass",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return error if password is greater than 100 characters",
			payload: types.LoginUserPayload{
				Email:    "test@example.com",
				Password: "fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rr := sendPayload(tc.payload, "/login", t)
			if rr.Code != tc.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, tc.wantStatus)
			}
		})
	}
}
