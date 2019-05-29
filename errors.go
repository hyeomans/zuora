package zuora

import (
	"fmt"
	"net/http"
	"strings"
)

type responseError struct {
	isTemporary bool
	message     string
}

func (r responseError) Temporary() bool {
	return r.isTemporary
}

func (r responseError) Error() string {
	return r.message
}

func isRetryableStatusCode(statusCode int) bool {
	return http.StatusRequestTimeout == statusCode ||
		http.StatusTooManyRequests == statusCode ||
		http.StatusInternalServerError == statusCode ||
		http.StatusServiceUnavailable == statusCode
}

//ErrorResponse could happen even when you get a 200 response from Zuora.
//The correct way to parse it, is to look at the Code inside reasons acording to Zuora docs.
type errorResponse struct {
	Success   bool   `json:"success"`
	ProcessID string `json:"processId"`
	Reasons   []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"reasons"`
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

func (e errorResponse) Temporary() bool {
	return isRetryableStatusCode(getStatus(e))
}

func (e errorResponse) Error() string {
	if len(e.Reasons) == 0 {
		return "there was an error on Zuora response but reasons array was empty."
	}

	allMessages := []string{}
	for i := 0; i < len(e.Reasons); i++ {
		currentCode := e.Reasons[i].Code
		currentMessage := e.Reasons[i].Message
		allMessages = append(allMessages, fmt.Sprintf("reason %v. Code: %v. Message: %v", i, currentCode, currentMessage))
	}

	return fmt.Sprintf("all errors from reasons: %v", strings.Join(allMessages, " -- "))
}

func getStatus(e errorResponse) int {
	if len(e.Reasons) == 0 {
		return http.StatusBadRequest
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

type tokenError struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Path      string `json:"path"`
	Reason    string `json:"reason"`
}

func (r tokenError) Temporary() bool {
	return false
}

func (r tokenError) Error() string {
	return fmt.Sprintf("could not generate token: %v - %v - %v - %v", r.Path, r.Reason, r.Status, r.Message)
}
