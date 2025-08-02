package server

import "http-job-que-system/types"

// The payload sent from client to our server for job creation
type JobRequest struct {
	Name string `json:"name"`
	Body string `json:"body"`
}

// Response to client after a successful job creation
type CreatedJobResponse struct {
	Message    string    `json:"message"`
	CreatedJob types.Job `json:"createdJob"`
}

// Structure or our error responses to clients
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}
