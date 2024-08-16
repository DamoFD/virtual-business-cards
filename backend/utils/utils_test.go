package utils

import (
    "net/http"
    "testing"
    "bytes"
    "encoding/json"
    "net/http/httptest"
)

func TestParseJSON(t *testing.T) {
    t.Run("invalid json", func(t *testing.T) {

        // Simulate an HTTP request with an invalid JSON body
        req, err := http.NewRequest("POST", "/", bytes.NewBufferString("{invalid json}"))
        if err != nil {
            t.Fatalf("failed to create request: %v", err)
        }

        payload := struct{
            Name string `json:"name"`
        }{}

        // Call the function with the invalid JSON
        err = ParseJSON(req, &payload)
        if err == nil {
            t.Fatal("expected error, got nil")
        }
    })

    t.Run("valid json", func(t *testing.T) {

        // Create a valid JSON payload
        validPayload := map[string]string{"name": "test"}
        validJSON, err := json.Marshal(validPayload)
        if err != nil {
            t.Fatalf("failed to marshal valid payload: %v", err)
        }

        // Simulate an HTTP request with a valid JSON body
        req, err := http.NewRequest("POST", "/", bytes.NewBuffer(validJSON))
        if err != nil {
            t.Fatalf("failed to create request: %v", err)
        }

        payload := struct {
            Name string `json:"name"`;
        }{}

        // Call the function with the valid JSON
        err = ParseJSON(req, &payload)
        if err != nil {
            t.Fatalf("expected no error, got %v", err)
        }

        // Verify that the payload was correctly populated
        if payload.Name != "test" {
            t.Errorf("expected name to be 'test', got %s", payload.Name)
        }
    })
}

func TestWriteJSON(t *testing.T) {
    t.Run("valid payload", func(t *testing.T) {

        // Arrange
        payload := map[string]string{"name": "test"}
        rr := httptest.NewRecorder()

        // Act
        err := WriteJSON(rr, http.StatusOK, payload)

        // Assert
        if err != nil {
            t.Fatalf("expected no error, got %v", err)
        }

        // Check the status code
        if status := rr.Code; status != http.StatusOK {
            t.Errorf("handler returned wrong status code: got %v want %v",
                status, http.StatusOK)
        }

        // Check the response body
        expectedBody, _ := json.Marshal(payload)
        if rr.Body.String() != string(expectedBody)+"\n" {
            t.Errorf("handler returned unexpected body: got %v want %v",
                rr.Body.String(), string(expectedBody)+"\n")
        }
    })

    t.Run("invalid payload", func(t *testing.T) {

        // Arrange
        payload := make(chan int) // Channels cannont be marshaled to JSON
        rr := httptest.NewRecorder()

        // Act
        err := WriteJSON(rr, http.StatusOK, payload)

        // Assert
        if err == nil {
            t.Fatal("expected error, got nil")
        }

        // Check that no body was written
        if rr.Body.Len() != 0 {
            t.Errorf("handler returned unexpected body: got %v want %v",
                rr.Body.String(), "")
        }

        // Check that the content type is still set correctly
        if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
            t.Errorf("handler returned unexpected content type: got %v want %v",
                contentType, "application/json")
        }
    })
}
