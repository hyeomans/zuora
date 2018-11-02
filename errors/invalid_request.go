package errors

import (
	"fmt"
	"net/http"
	"strings"
)

type InvalidRequestError struct {
	jsonInvalidRequestError *jsonInvalidRequestError
}

type jsonInvalidRequestError struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Path      string `json:"path"`
}

func newInvalidRequestError(jsonInvalidRequestError *jsonInvalidRequestError) *InvalidRequestError {
	return &InvalidRequestError{
		jsonInvalidRequestError: jsonInvalidRequestError,
	}
}

func (e *InvalidRequestError) Temporary() bool {
	if e.jsonInvalidRequestError.Status == http.StatusTooManyRequests {
		return true
	}

	return false
}

func (e *InvalidRequestError) Error() string {
	var finalMessage []string
	if e.jsonInvalidRequestError.Path != "" {
		finalMessage = append(finalMessage, fmt.Sprintf("path: %v", e.jsonInvalidRequestError.Path))
	}
	if e.jsonInvalidRequestError.Timestamp != "" {
		finalMessage = append(finalMessage, fmt.Sprintf("timestamp: %v", e.jsonInvalidRequestError.Timestamp))
	}

	if len(finalMessage) > 0 {
		return fmt.Sprintf("zuora returned error - statusCode: %v message: %v %v", e.jsonInvalidRequestError.Status, e.jsonInvalidRequestError.Error, strings.Join(finalMessage, " "))
	}
	return fmt.Sprintf("zuora returned error - statusCode: %v message: %v", e.jsonInvalidRequestError.Status, e.jsonInvalidRequestError.Error)
}
