// Package server provides the HTTP server setup and startup logic
// for the job queue system. It handles initialization and bootstrapping
// of the application’s HTTP services.
package server

import (
	"encoding/json"
	"fmt"
	"http-job-que-system/logger"
	"net/http"
	"slices"
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

	http.HandleFunc("/", jobHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Log.Fatal("Server failed to start:", err)
	}
}

// Main entry point to our job que system
func jobHandler(w http.ResponseWriter, r *http.Request) {

	// Reject invalid http methods
	notAllowedMethods := []string{
		http.MethodGet,
		http.MethodPatch,
		http.MethodPut,
		http.MethodDelete,
	}
	if slices.Contains(notAllowedMethods, r.Method) {
		responseStatus := http.StatusMethodNotAllowed
		responseBody := fmt.Sprintf("%d %s", responseStatus, http.StatusText(responseStatus))
		http.Error(w, responseBody, responseStatus)

		logMessage := fmt.Sprintf(
			"jobHandler: Disallowed HTTP method '%s' for path '%s' from client '%s'. Responded with %d.",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			responseStatus,
		)
		logger.Log.Println(logMessage)
		fmt.Println(logMessage)

		return
	}

	if r.Method == http.MethodPost {
		var clientMsg Message

		// Parse response body
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&clientMsg)
		if err != nil {
			errMsg := "Error decoding JSON from request body"
			fmt.Println(errMsg, err)
			logger.Log.Println(errMsg)
			responseStatus := http.StatusBadRequest
			responseBody := fmt.Sprintf("%d %s", responseStatus, http.StatusText(responseStatus))
			http.Error(w, responseBody, responseStatus)
			return
		}

		// Parse successful, process client job

		// Respond to client post job processing
		responseStatus := http.StatusCreated
		responseBody := fmt.Sprintf(
			"%d %s %#v",
			responseStatus,
			http.StatusText(responseStatus),
			clientMsg,
		)
		w.WriteHeader(responseStatus)
		_, err = fmt.Fprint(w, responseBody)
		if err != nil {
			errMsg := "Error writing to response body"
			fmt.Println(errMsg, err)
			logger.Log.Println(errMsg)
		}

		fmt.Println(responseBody)
		return
	}

	// DEFAULT: return forbidden status
	responseStatus := http.StatusForbidden
	responseBody := fmt.Sprintf("%d %s", responseStatus, http.StatusText(responseStatus))
	http.Error(w, responseBody, responseStatus)
}
