package app_errors

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	Code    int
	Message string
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func NewNotFoundError(message string) error {
	return &HttpError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError(message string) error {
	return &HttpError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewBadRequestError(message string) error {
	return &HttpError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}
