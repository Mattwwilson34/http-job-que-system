package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestFooHandler(t *testing.T) {
	var jobHandlerUrl = "http://localhost:8080"

	// Should return method not allowed for GET, PATCH, PUT, DELETE
	// GET
	resp, err := http.Get(jobHandlerUrl)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf(
			"Expected status %d for GET, got %d",
			http.StatusMethodNotAllowed,
			resp.StatusCode,
		)
	}

	// PATCH
	req, err := http.NewRequest(http.MethodPatch, jobHandlerUrl, nil)
	if err != nil {
		t.Fatalf("Failed to create Patch request %v", err)
	}
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Patch request failed %v", err)
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf(
			"Expected status %d for PATCH, got %d",
			http.StatusMethodNotAllowed,
			resp.StatusCode,
		)
	}

	// PUT
	req, err = http.NewRequest(http.MethodPut, jobHandlerUrl, nil)
	if err != nil {
		t.Fatalf("Failed to create Put request %v", err)
	}
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Put request failed %v", err)
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf(
			"Expected status %d for PUT, got %d",
			http.StatusMethodNotAllowed,
			resp.StatusCode,
		)
	}

	// DELETE
	req, err = http.NewRequest(http.MethodDelete, jobHandlerUrl, nil)
	if err != nil {
		t.Fatalf("Failed to create Delete request %v", err)
	}
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Delete request failed %v", err)
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf(
			"Expected status %d for DELETE, got %d",
			http.StatusMethodNotAllowed,
			resp.StatusCode,
		)
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
