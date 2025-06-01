// Package server provides the HTTP server setup and startup logic
// for the job queue system. It handles initialization and bootstrapping
// of the application’s HTTP services.
package server

import (
	"fmt"
	"http-job-que-system/logger"
	"net/http"
)

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

	http.HandleFunc("/", HandleJobRequest)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Log.Fatal("Server failed to start:", err)
	}
}
