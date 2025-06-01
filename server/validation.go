package server

import (
	"fmt"
	"net/http"
)

// Validate a Job Request
func ValidateJobRequest(job *JobRequest) error {
	if job.Name == "" {
		return fmt.Errorf("name is a required job field")
	}
	if job.Body == "" {
		return fmt.Errorf("body is a required job field")
	}
	return nil
}

// Validate that the request is a POST request
func ValidateHttpMethod(r *http.Request) error {
	if r.Method != http.MethodPost {
		return fmt.Errorf("method %s not allowed", r.Method)
	}

	return nil
}
