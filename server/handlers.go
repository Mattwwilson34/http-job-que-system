package server

import (
	"encoding/json"
	"fmt"
	"http-job-que-system/utils"
	"net/http"
	"time"
)

// Main entry point to our job que system
func HandleJobRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handleJobCreation(w, r)
	default:
		RejectRequest(
			w,
			r,
			http.StatusMethodNotAllowed,
			"handleMethodNotAllowed",
			http.StatusText(http.StatusMethodNotAllowed),
		)
	}

}

// Attempts to create a job from the client request, if it fails at any step
// then the client request is rejected
func handleJobCreation(w http.ResponseWriter, r *http.Request) bool {
	var jobRequest JobRequest

	// Parse response body
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&jobRequest)
	if err != nil {
		errorMsg := "Error decoding JSON Job from request body"
		RejectRequest(w, r, http.StatusBadRequest, "jobHandler", errorMsg)
		return false
	}

	// Validate JobRequest payload
	err = ValidateJobRequest(&jobRequest)
	if err != nil {
		errorMsg := fmt.Sprintf("Job request is not valid, %s", err)
		RejectRequest(w, r, http.StatusBadRequest, "jobHandler", errorMsg)
		return false
	}

	// Generate Job Id
	jobId, err := uuid.GenerateUUID()
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to create Job UUID, %s", err)
		RejectRequest(w, r, http.StatusInternalServerError, "jobHandler", errorMsg)
		return false
	}

	// Get Created DateTime
	createdDateTime := time.Now().UTC().Format(time.RFC3339)

	// Create Job
	job := Job{jobId, jobRequest.Name, jobRequest.Body, createdDateTime}

	SendJobCreatedResponse(w, job)
	return true
}
