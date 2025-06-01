package server

import (
	"encoding/json"
	"fmt"
	"http-job-que-system/logger"
	"net/http"
)

// Respond to client with successful job creation message. Return true if the
// response was sent successfully and false on failure.
func SendJobCreatedResponse(w http.ResponseWriter, clientJob Job) bool {
	createdJobResponse := CreatedJobResponse{"Job creation successful", clientJob}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err := json.NewEncoder(w).Encode(createdJobResponse)
	if err != nil {
		errMsg := "Error writing to response body"
		fmt.Println(errMsg, err)
		logger.Log.Println(errMsg)
		return false
	}

	logData, _ := json.Marshal(createdJobResponse)
	fmt.Printf("Response sent: %s\n", string(logData))
	return true
}

// Reject client HTTP request with custom error message
func RejectRequest(
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
