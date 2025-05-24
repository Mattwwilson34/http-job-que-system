package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestFooHandler(t *testing.T) {
	var jobHandlerUrl = "http://localhost:8080"
	var client = &http.Client{}

	// Reject all non-Post http methods
	notAllowedMethods := []string{
		http.MethodGet,
		http.MethodPatch,
		http.MethodPut,
		http.MethodDelete,
	}
	for _, method := range notAllowedMethods {
		req, err := http.NewRequest(method, jobHandlerUrl, nil)
		if err != nil {
			t.Fatalf("Failed to create %s request %v", method, err)
		}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}
		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf(
				"Expected status %d for %s, got %d",
				http.StatusMethodNotAllowed,
				method,
				resp.StatusCode,
			)
		}
	}

	// Why when we declare our json body outside the loop do we only get 1 post
	// response body
	for i := range 10 {
		m := Message{i, "Alice", "Hello", 1294706395881547000}
		b, err := json.Marshal(m)
		if err != nil {
			t.Fatalf("Failed to marshal JSON")
		}

		buf := bytes.NewBuffer(b)

		resp, err := http.Post(jobHandlerUrl, "application/json", buf)
		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Unexpected status: %d", resp.StatusCode)
		}

		resp.Body.Close()
	}
}
