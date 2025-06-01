package server

import "fmt"

// Extended error interface with isClientError boolean
type ClientError interface {
	error
	IsClientError() bool
}

// Extended error interface with isServerError boolean
type ServerError interface {
	error
	IsServerError() bool
}

// Custom client validation error struct
type ValidationError struct {
	Field   string
	Message string
}

// Custom server error struct
type InternalError struct {
	Operation string
	Cause     error
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error in field %s: %s", e.Field, e.Message)
}

func (e InternalError) Error() string {
	return fmt.Sprintf("internal error during %s: %v", e.Operation, e.Cause)
}

func (e ValidationError) IsClientError() bool {
	return true
}

func (e InternalError) IsServerError() bool {
	return true
}
