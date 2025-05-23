package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestFooHandler(t *testing.T) {

	// Why when we declare our json body outside the loop do we only get 1 post
	// response body

	for i := range 10 {
		m := Message{i, "Alice", "Hello", 1294706395881547000}
		b, err := json.Marshal(m)
		if err != nil {
			t.Fatalf("Failed to marshal JSON")
		}

		buf := bytes.NewBuffer(b)

		resp, err := http.Post("http://localhost:8080/foo", "application/json", buf)
		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Unexpected status: %d", resp.StatusCode)
		}

		resp.Body.Close()
	}
}
