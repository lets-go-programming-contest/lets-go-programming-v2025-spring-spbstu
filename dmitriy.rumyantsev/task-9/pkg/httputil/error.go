package httputil

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ErrorResponse represents a structure for error responses
type ErrorResponse struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
	Details  string `json:"details,omitempty"`
}

// Error implements the error interface
func (e ErrorResponse) Error() string {
	return fmt.Sprintf("HTTP: %d, Code: %d, Message: %s, Details: %s",
		e.HTTPCode, e.Code, e.Message, e.Details)
}

// JSONError sends an error to the client in JSON format
func JSONError(w http.ResponseWriter, e ErrorResponse) {
	data := struct {
		Error ErrorResponse `json:"error"`
	}{e}

	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HTTPCode)
	w.Write(b)
}

// NewBadRequestError creates an error with code 400 Bad Request
func NewBadRequestError(message, details string) ErrorResponse {
	return ErrorResponse{
		HTTPCode: http.StatusBadRequest,
		Code:     400001,
		Message:  message,
		Details:  details,
	}
}

// NewNotFoundError creates an error with code 404 Not Found
func NewNotFoundError(message, details string) ErrorResponse {
	return ErrorResponse{
		HTTPCode: http.StatusNotFound,
		Code:     404001,
		Message:  message,
		Details:  details,
	}
}

// NewInternalServerError creates an error with code 500 Internal Server Error
func NewInternalServerError(details string) ErrorResponse {
	return ErrorResponse{
		HTTPCode: http.StatusInternalServerError,
		Code:     500001,
		Message:  "Internal Server Error",
		Details:  details,
	}
}
