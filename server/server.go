// Package server provides the HTTP server setup and startup logic
// for the job queue system. It handles initialization and bootstrapping
// of the application’s HTTP services.
package server

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"http-job-que-system/logger"
	"net/http"
	"time"
)

type Job struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Body            string `json:"body"`
	CreatedDateTime string `json:"createdDateTime"`
}

type JobRequest struct {
	Name string `json:"name"`
	Body string `json:"body"`
}

type CreatedJobResponse struct {
	Message    string `json:"message"`
	CreatedJob Job    `json:"createdJob"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
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
	err := validateHttpMethod(r)
	if err != nil {
		errorMsg := fmt.Sprintf("Http Method %s not allowed", r.Method)
		rejectRequest(w, r, http.StatusMethodNotAllowed, "jobHandler", errorMsg)
		return
	}

	if r.Method == http.MethodPost {
		var jobRequest JobRequest

		// Parse response body
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&jobRequest)
		if err != nil {
			errorMsg := "Error decoding JSON Job from request body"
			rejectRequest(w, r, http.StatusBadRequest, "jobHandler", errorMsg)
			return
		}

		// Validate JobRequest payload
		err = validateJobRequest(&jobRequest)
		if err != nil {
			errorMsg := fmt.Sprintf("Job request is not valid, %s", err)
			rejectRequest(w, r, http.StatusBadRequest, "jobHandler", errorMsg)
			return
		}

		// Generate Job Id
		jobId, err := GenerateRandomID()
		if err != nil {
			errorMsg := fmt.Sprintf("Failed to create Job UUID, %s", err)
			rejectRequest(w, r, http.StatusInternalServerError, "jobHandler", errorMsg)
			return
		}

		// Get Created DateTime
		createdDateTime := time.Now().UTC().Format(time.RFC3339)

		// Create Job
		job := Job{jobId, jobRequest.Name, jobRequest.Body, createdDateTime}

		sendJobCreatedResponse(w, job)
		return
	}

	// DEFAULT: return forbidden status
	responseStatus := http.StatusForbidden
	responseBody := fmt.Sprintf("%d %s", responseStatus, http.StatusText(responseStatus))
	http.Error(w, responseBody, responseStatus)
}

// Validate a Job
func validateJobRequest(job *JobRequest) error {
	if job.Name == "" {
		return fmt.Errorf("name is a required job field")
	}
	if job.Body == "" {
		return fmt.Errorf("body is a required job field")
	}
	return nil
}

// Throw error if not  POST request
func validateHttpMethod(r *http.Request) error {
	if r.Method != http.MethodPost {
		return fmt.Errorf("method %s not allowed", r.Method)
	}

	return nil
}

// Respond to users with successful job creation message
func sendJobCreatedResponse(w http.ResponseWriter, clientJob Job) {
	createdJobResponse := CreatedJobResponse{"Job creation successful", clientJob}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err := json.NewEncoder(w).Encode(createdJobResponse)
	if err != nil {
		errMsg := "Error writing to response body"
		fmt.Println(errMsg, err)
		logger.Log.Println(errMsg)
	}

	logData, _ := json.Marshal(createdJobResponse)
	fmt.Printf("Response sent: %s\n", string(logData))

}

// Reject HTTP request if Method is not POST
func rejectRequest(
	w http.ResponseWriter,
	r *http.Request,
	responseStatus int,
	caller string,
	errorMessage string,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)

	errorResponse := ErrorResponse{
		Error:   http.StatusText(responseStatus),
		Message: errorMessage,
		Status:  http.StatusMethodNotAllowed,
	}

	err := json.NewEncoder(w).Encode(errorResponse)
	if err != nil {
		logger.Log.Println("rejectRequest: Error encoding ErrorResponse to JSON", err)
	}

	logInput := logger.HttpLogMsg{
		FuncName:   caller,
		Method:     r.Method,
		UrlPath:    r.URL.Path,
		RemoteAddr: r.RemoteAddr,
		Status:     responseStatus,
		Message:    errorMessage,
	}

	logger.Log.Println(logInput)
	fmt.Println(logInput)
}

// Generate a random uuid
func GenerateRandomID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
