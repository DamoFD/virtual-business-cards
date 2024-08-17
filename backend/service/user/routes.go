/*
Package user contains user routes and methods.
It contains a struct Handler and a method RegisterRoutes() that registers the user routes and methods.
It contains a method handleLogin() that handles the login route.
It contains a method handleRegister() that handles the register route.
*/
package user

import (
	"fmt"
	"net/http"
	"time"

	validator "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/DamoFD/virtual-business/config"
	"github.com/DamoFD/virtual-business/types"
	"github.com/DamoFD/virtual-business/utils"
)

// Handler is a struct that contains a user store.
// It contains a method RegisterRoutes() that registers the user routes and methods.
// It contains a method handleLogin() that handles the login route.
// It contains a method handleRegister() that handles the register route.
type Handler struct {
	store types.UserStore // user store
	auth  types.Auth
}

// NewHandler creates a new user handler.
// It takes a user store as a parameter.
// It returns a pointer to the Handler struct.
func NewHandler(store types.UserStore, auth types.Auth) *Handler {
	return &Handler{
		store: store,
		auth:  auth,
	}
}

// RegisterRoutes registers the user routes and methods.
// It takes a router as a parameter.
func (h *Handler) RegisterRoutes(router *mux.Router, middleware types.Middleware) {
	router.Handle("/login", middleware.RateLimit(10, time.Minute)(http.HandlerFunc(h.handleLogin))).Methods("POST")
	router.Handle("/register", middleware.RateLimit(10, time.Minute)(http.HandlerFunc(h.handleRegister))).Methods("POST")
}

// handleLogin handles the login route.
// It takes a http.ResponseWriter and a *http.Request as parameters.
// It returns an error if the login fails.
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload

	// Parse the request body into the LoginUserPayload struct
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate the LoginUserPayload struct
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// Check if the user exists by email
	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid credentials"))
		return
	}

	// Compare the plain password with the hashed password
	if !h.auth.ComparePassword(u.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid credentials"))
		return
	}

	// Create a JWT token
	secret := []byte(config.Envs.JWTSecret)
	token, err := h.auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Return the JWT token
	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

// handleRegister handles the register route.
// It takes a http.ResponseWriter and a *http.Request as parameters.
// It returns an error if the registration fails.
// User must not exist
// Passwords must match
// Password must be hashed
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	// get JSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// check if the user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("user already exists"))
		return
	}

	// confirm the password
	confirmed := h.auth.ConfirmPassword(payload.Password, payload.ConfirmPassword)
	if !confirmed {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("passwords do not match"))
		return
	}

	// hash the password
	hashedPassword, err := h.auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// if it doesn't exist, create the user
	user := types.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashedPassword,
	}
	err = h.store.CreateUser(user)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// return StatusCreated
	utils.WriteJSON(w, http.StatusCreated, nil)
}
