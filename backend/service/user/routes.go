/*
Package user contains user routes and methods.
It contains a struct Handler and a method RegisterRoutes() that registers the user routes and methods.
It contains a method handleLogin() that handles the login route.
It contains a method handleRegister() that handles the register route.
*/
package user

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/DamoFD/virtual-business/types"
	"github.com/DamoFD/virtual-business/utils"
)

// Handler is a struct that contains a user store.
// It contains a method RegisterRoutes() that registers the user routes and methods.
// It contains a method handleLogin() that handles the login route.
// It contains a method handleRegister() that handles the register route.
type Handler struct {
	store types.UserStore // user store
}

// NewHandler creates a new user handler.
// It takes a user store as a parameter.
// It returns a pointer to the Handler struct.
func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

// RegisterRoutes registers the user routes and methods.
// It takes a router as a parameter.
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

// handleLogin handles the login route.
// It takes a http.ResponseWriter and a *http.Request as parameters.
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
}

// handleRegister handles the register route.
// It takes a http.ResponseWriter and a *http.Request as parameters.
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//
}
