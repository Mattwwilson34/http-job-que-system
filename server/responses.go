package server

import (
	"encoding/json"
	"fmt"
	"http-job-que-system/logger"
	"net/http"
)

// Respond to client with successful job creation message. Returns an error if json encoding fails
func SendJobCreatedResponse(w http.ResponseWriter, clientJob Job) error {
	createdJobResponse := CreatedJobResponse{"Job creation successful", clientJob}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err := json.NewEncoder(w).Encode(createdJobResponse)
	if err != nil {
		return fmt.Errorf("Error writing to response body, %s", err.Error())
	}

	logData, _ := json.Marshal(createdJobResponse)
	fmt.Printf("Response sent: %s\n", string(logData))
	return nil
}

// Reject client HTTP request with custom error message
func RejectRequest(
	w http.ResponseWriter,
	r *http.Request,
	responseStatus int,
	caller string,
	clientMessage string,
	errorForLogs error,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)

	errorResponse := ErrorResponse{
		Error:   http.StatusText(responseStatus),
		Message: clientMessage,
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
		Message:    errorForLogs.Error(),
	}

	logger.Log.Println(logInput)
	fmt.Println(logInput)
}
