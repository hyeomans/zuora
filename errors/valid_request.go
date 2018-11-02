package errors

import (
	"fmt"
	"net/http"
	"strings"
)

type ValidRequestError struct {
	Message string
	Status  int
}

type jsonValidRequestError struct {
	Success   bool                     `json:"success"`
	ProcessID string                   `json:"processId"`
	Reasons   []jsonReasonRequestError `json:"reasons"`
	Message   string
	Status    int
}

type jsonReasonRequestError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func newValidRequestError(jsonValidRequestError *jsonValidRequestError) *ValidRequestError {
	return &ValidRequestError{
		Status:  generateStatus(jsonValidRequestError),
		Message: generateErrorMsg(jsonValidRequestError),
	}
}

func (e *ValidRequestError) Error() string {
	return fmt.Sprintf("zuora returned an error (status: %v): %v", e.Status, e.Message)
}

func (e *ValidRequestError) Temporary() bool {
	switch e.Status {
	case http.StatusTooManyRequests, http.StatusLocked, http.StatusInternalServerError:
		return true
	}

	return false
}

const (
	//Unknown
	unknown = 0

	//Permission or access denied
	accessDenied = 10

	//Authentication failed
	authFailed = 11

	//Invalid format or value
	invalidFormat = 20

	//Unknown field in request
	unknownField = 21

	//Missing required field
	requiredField = 22

	//Rule restriction
	ruleRestriction = 30

	//Not found
	notFound = 40

	//Locking contention
	lockingContention = 50

	//Internal error
	internalError = 60

	//Request exceeded limit
	requestExceeded = 70

	//Malformed request
	malformedRequest = 90

	//Extension error
	extensionError = 99
)

func generateErrorMsg(e *jsonValidRequestError) string {
	if len(e.Reasons) == 0 {
		return e.Message
	}

	var errorMessage []string

	for i := 0; i < len(e.Reasons); i++ {
		errorMessage = append(errorMessage, fmt.Sprintf("code: %v message: %v", e.Reasons[i].Code, strings.ToLower(e.Reasons[i].Message)))
	}

	return fmt.Sprint(strings.Join(errorMessage, " || "))
}

func generateStatus(e *jsonValidRequestError) int {
	if len(e.Reasons) == 0 {
		return e.Status
	}

	allCodes := []int{}
	for i := 0; i < len(e.Reasons); i++ {
		currentCode := e.Reasons[i].Code
		category := currentCode % 100
		allCodes = append(allCodes, category)
	}

	var isRequestExceeded bool
	var isAccessDenied bool
	var isAuthDenied bool
	var isNotFound bool
	var isStatusLocked bool
	var isInternalServerError bool

	for i := 0; i < len(allCodes); i++ {
		currentCode := allCodes[i]
		switch currentCode {
		case requestExceeded:
			isRequestExceeded = true
		case accessDenied:
			isAccessDenied = true
		case authFailed:
			isAuthDenied = true
		case notFound:
			isNotFound = true
		case lockingContention:
			isStatusLocked = true
		case internalError:
			isInternalServerError = true
		}
	}

	if isRequestExceeded {
		return http.StatusTooManyRequests
	} else if isAuthDenied {
		return http.StatusUnauthorized
	} else if isAccessDenied {
		return http.StatusForbidden
	} else if isNotFound {
		return http.StatusNotFound
	} else if isStatusLocked {
		return http.StatusLocked
	} else if isInternalServerError {
		return http.StatusInternalServerError
	}

	return http.StatusBadRequest
}
