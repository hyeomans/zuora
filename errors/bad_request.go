package errors

import (
	"fmt"
)

type BadRequestError struct {
	err error
}

func newBadRequestError(err error) *BadRequestError {
	return &BadRequestError{err: err}
}

func (e *BadRequestError) Temporary() bool {
	return false
}

func (e *BadRequestError) Error() string {
	return fmt.Sprintf("%v", e.err)
}
