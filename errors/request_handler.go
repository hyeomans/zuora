package errors

import (
	"encoding/json"
	"fmt"
)

type RequestHandler struct{}

func (e RequestHandler) InvalidResponse(body []byte, statusCode int) error {
	jsonError := &jsonInvalidRequestError{}

	if err := json.Unmarshal(body, jsonError); err != nil {
		return newInvalidRequestError(&jsonInvalidRequestError{Status: statusCode, Error: string(body)})
	}

	if jsonError.Error == "" {
		switch statusCode {
		case 401:
			return newInvalidRequestError(&jsonInvalidRequestError{Status: statusCode, Error: "Authentication error"})
		case 404:
			return newInvalidRequestError(&jsonInvalidRequestError{Status: statusCode, Error: "object not found"})
		default:
			return newInvalidRequestError(&jsonInvalidRequestError{Status: statusCode, Error: string(body)})
		}
	}

	return newInvalidRequestError(jsonError)
}

func (e RequestHandler) ValidRequestError(body []byte, statusCode int) error {
	jsonError := &jsonValidRequestError{}

	if err := json.Unmarshal(body, jsonError); err != nil {
		jsonError.Status = statusCode
		jsonError.Message = fmt.Sprint(err)
		return newValidRequestError(jsonError)
	}

	return newValidRequestError(jsonError)
}

func (e RequestHandler) BadRequest(err error) error {
	return newBadRequestError(err)
}
