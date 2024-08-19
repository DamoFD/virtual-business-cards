package user

import (
	"net/http"
	"testing"

	"github.com/DamoFD/virtual-business/types"
)

func TestValidateRegisterPayload(t *testing.T) {

	testCases := []struct {
		name       string
		payload    types.RegisterUserPayload
		wantStatus int
	}{
		{
			name: "Should return error if name is empty",
			payload: types.RegisterUserPayload{
				Name:            "",
				Email:           "test@example.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return error if email is empty",
			payload: types.RegisterUserPayload{
				Name:            "test",
				Email:           "",
				Password:        "password",
				ConfirmPassword: "password",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return error if password is empty",
			payload: types.RegisterUserPayload{
				Name:            "test",
				Email:           "test@example.com",
				Password:        "",
				ConfirmPassword: "password",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return error if confirm password is empty",
			payload: types.RegisterUserPayload{
				Name:            "test",
				Email:           "test@example.com",
				Password:        "password",
				ConfirmPassword: "",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return error if name is less than 2 characters",
			payload: types.RegisterUserPayload{
				Name:            "t",
				Email:           "test@example.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return error if name is greater than 100 characters",
			payload: types.RegisterUserPayload{
				Name:            "fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
				Email:           "test@example.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return error if email is invalid",
			payload: types.RegisterUserPayload{
				Name:            "test",
				Email:           "test",
				Password:        "password",
				ConfirmPassword: "password",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return error if email is greater than 320 characters",
			payload: types.RegisterUserPayload{
				Name:            "test",
				Email:           "fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff@gmail.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return error if password is less than 6 characters",
			payload: types.RegisterUserPayload{
				Name:            "test",
				Email:           "test@example.com",
				Password:        "pass",
				ConfirmPassword: "password",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return error if password is greater than 100 characters",
			payload: types.RegisterUserPayload{
				Name:            "test",
				Email:           "test@example.com",
				Password:        "fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
				ConfirmPassword: "password",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return error if confirm password is less than 6 characters",
			payload: types.RegisterUserPayload{
				Name:            "test",
				Email:           "test@example.com",
				Password:        "password",
				ConfirmPassword: "pass",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return error if password is greater than 100 characters",
			payload: types.RegisterUserPayload{
				Name:            "test",
				Email:           "test@example.com",
				Password:        "password",
				ConfirmPassword: "fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rr := sendPayload(tc.payload, "/register", t)
			if rr.Code != tc.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, tc.wantStatus)
			}
		})
	}
}
