package server

import (
	"encoding/json"
	"errors"
	"http-job-que-system/logger"
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
			errors.New(http.StatusText(http.StatusMethodNotAllowed)),
		)
	}

}

// Attempts to create a job from the client request, if it fails at any step
// then the client request is rejected
func handleJobCreation(w http.ResponseWriter, r *http.Request) {
	job, err := createJobFromRequest(r)

	if err != nil {
		var clientErr ClientError
		var serverErr ServerError

		switch {
		case errors.As(err, &clientErr):
			RejectRequest(w, r, http.StatusBadRequest, "handleJobCreation", err.Error(), err)

		case errors.As(err, &serverErr):
			RejectRequest(
				w,
				r,
				http.StatusInternalServerError,
				"handleJobCreation",
				http.StatusText(http.StatusInternalServerError),
				err,
			)

		default:
			RejectRequest(
				w,
				r,
				http.StatusInternalServerError,
				"handleJobCreation",
				"Unknown error",
				err,
			)
		}
		return
	}

	err = SendJobCreatedResponse(w, job)
	if err != nil {
		logger.Log.Println(
			"SendJobCreatedResponse: failed to send successful job creation response to client",
			err,
		)
	}
}

// Create a job from a client request
func createJobFromRequest(r *http.Request) (Job, error) {
	var jobRequest JobRequest

	// Parse request
	if err := json.NewDecoder(r.Body).Decode(&jobRequest); err != nil {
		return Job{}, ValidationError{
			Field:   "request_body",
			Message: "invalid JSON: " + err.Error(),
		}
	}

	// Validate parsed request body is valid
	if err := ValidateJobRequest(&jobRequest); err != nil {
		return Job{}, ValidationError{
			Field:   "payload",
			Message: "validation failed: " + err.Error(),
		}
	}

	// Create job
	job, err := createJob(&jobRequest)
	if err != nil {
		return Job{}, InternalError{
			Operation: "job_creation",
			Cause:     err,
		}
	}

	logger.Log.Printf("createJobFromRequest: job (%s) created with id:(%s)", job.Name, job.Id)
	return job, nil
}

// Create and return a job or return an error
func createJob(req *JobRequest) (Job, error) {
	jobID, err := utils.GenerateUUID()
	if err != nil {
		return Job{}, err
	}

	return Job{
		Id:              jobID,
		Name:            req.Name,
		Body:            req.Body,
		CreatedDateTime: time.Now().UTC().Format(time.RFC3339),
	}, nil
}
