package user

import (
    "net/http"

    "github.com/gorilla/mux"

    "github.com/DamoFD/virtual-business/types"
    "github.com/DamoFD/virtual-business/utils"
)

type Handler struct {
    store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
    return &Handler{
        store: store,
    }
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/login", h.handleLogin).Methods("POST")
    router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
    var payload types.LoginUserPayload
    if err := utils.ParseJSON(r, &payload); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }
}

func (h *Handler) handleRegister (w http.ResponseWriter, r *http.Request) {
    //
}
