// Package server provides the HTTP server setup and startup logic
// for the job queue system. It handles initialization and bootstrapping
// of the application’s HTTP services.
package server

import (
	"encoding/json"
	"fmt"
	"http-job-que-system/logger"
	"net/http"
)

type Message struct {
	Id   int
	Name string
	Body string
	Time int64
}

// Start launches the HTTP server on port 8080.
// It logs a startup message using the global logger and begins listening
// for incoming HTTP requests. If the server fails to start, it logs the error
// and exits the application.
//
// This function should be called only after the logger has been initialized.
func Start() {
	startupMsg := "Server listening on port 8080"

	logger.Log.Println(startupMsg)
	fmt.Println(startupMsg)

	http.HandleFunc("/foo", fooHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Log.Fatal("Server failed to start:", err)
	}
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	var m Message
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&m)
	if err != nil {
		fmt.Println("Decode error:", err) // Add this before Fatal
		logger.Log.Fatal("Error decoding request body json")
	}
	fmt.Println("Message:", m)
}
