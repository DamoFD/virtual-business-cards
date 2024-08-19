/*
Package utils contains utility functions for the API server.
It contains a struct utils and a method ParseJSON() that parses a JSON request body into a struct
and a method WriteJSON() that writes a JSON response with a status code
and a method WriteError() that writes an error response with a status code
*/
package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Validate is a validator instance
var Validate = validator.New()

func init() {
	Validate.RegisterValidation("slug", validateSlug)
}

// ParseJSON parses a JSON request body into a struct
// It takes a *http.Request and a pointer to a struct as arguments.
// and returns an error if the parsing fails.
func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

// WriteJSON writes a JSON response with a status code
// It takes a http.ResponseWriter, a status code, and a payload as arguments.
// and returns an error if the writing fails.
func WriteJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(payload)
}

// WriteError writes an error response with a status code
// It takes a http.ResponseWriter, a status code, and an error as arguments.
// and returns an error if the writing fails.
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func validateSlug(fl validator.FieldLevel) bool {
	slug := fl.Field().String()

	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

	return slugRegex.MatchString(slug)
}
