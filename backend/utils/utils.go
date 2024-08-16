package utils

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func ParseJSON(r *http.Request, payload any) error {
    if r.Body == nil {
        return fmt.Errorf("missing request body")
    }

    return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, payload any) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)

    return json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, status int, err error) {
    WriteJSON(w, status, map[string]string{"error": err.Error()})
}
