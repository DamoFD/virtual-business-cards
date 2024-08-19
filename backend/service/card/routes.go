package card

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/DamoFD/virtual-business/types"
	"github.com/DamoFD/virtual-business/utils"
)

// Handler is a struct that contains a user store.
// It contains a method RegisterRoutes() that registers the card routes and methods.
type Handler struct {
	store types.CardStore // user store
	auth  types.Auth      // auth service
}

// NewHandler creates a new user handler.
// It takes a user store and an auth service as parameters.
// It returns a pointer to the Handler struct.
func NewHandler(store types.CardStore, auth types.Auth) *Handler {
	return &Handler{
		store: store,
		auth:  auth,
	}
}

// RegisterRoutes registers the user routes and methods.
// It takes a router and a middleware as parameters.
func (h *Handler) RegisterRoutes(router *mux.Router, middleware types.Middleware) {

	// Public routes
	router.Handle("/{slug}", middleware.RateLimit(100, time.Minute)(http.HandlerFunc(h.handleGetCard))).Methods("GET")

	// Authenticated route group
	authRouter := router.PathPrefix("/").Subrouter()
	authRouter.Use(middleware.Auth())
	authRouter.Use(middleware.RateLimit(100, time.Minute))

	// Protected routes
	authRouter.Handle("/cards/@me", http.HandlerFunc(h.handleGetMyCards)).Methods("GET")
	authRouter.Handle("/cards", http.HandlerFunc(h.handleCreateCard)).Methods("POST")
}

func (h *Handler) handleGetCard(w http.ResponseWriter, r *http.Request) {

	// Extract the slug from the URL
	vars := mux.Vars(r)
	slug := vars["slug"]

	// Get the card by slug
	card, err := h.store.GetCardBySlug(slug)
	if err != nil {
		if err.Error() == "card not found" {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("card not found"))
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, card); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to write response"))
	}
}

func (h *Handler) handleGetMyCards(w http.ResponseWriter, r *http.Request) {

	// Retrieve the user ID from the Redis session
	sessionID, err := r.Cookie("session_id")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("session not found"))
		return
	}

	sessionData, err := h.auth.GetSession(r.Context(), sessionID.Value)
	if err != nil || sessionData.UserID == 0 {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("session not found"))
		return
	}

	cards, err := h.store.GetCardsByUserID(sessionData.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, cards); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to write response"))
	}
}

func (h *Handler) handleCreateCard(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.CardPayload
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

	// check if the slug is used
	_, err := h.store.GetCardBySlug(payload.Slug)
	if err == nil {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("Slug is taken"))
		return
	}

	// Retrieve the user ID from the Redis session
	sessionID, err := r.Cookie("session_id")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("session not found"))
		return
	}

	sessionData, err := h.auth.GetSession(r.Context(), sessionID.Value)
	if err != nil || sessionData.UserID == 0 {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("session not found"))
		return
	}

	// Create the card
	card, err := h.store.CreateCard(r.Context(), types.Card{
		Name:   payload.Name,
		Slug:   payload.Slug,
		UserID: sessionData.UserID,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Create the associated CardItemFields
	for _, field := range payload.CardItems {
		cardItemField := types.CardItemField{
			CardID:     card.ID,
			CardItemID: field.CardItemID,
			Name:       field.Name,
			Value:      field.Value,
		}

		if err := h.store.CreateCardItemField(r.Context(), cardItemField); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create card item field: %v", err))
			return
		}
	}

	// Return a success response
	utils.WriteJSON(w, http.StatusCreated, card)
}
