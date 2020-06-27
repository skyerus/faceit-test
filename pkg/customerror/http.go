package customerror

import (
	"net/http"
)

// Error ...
type Error interface {
	OriginalError() error
	Code() int
	Message() string
}

// HTTPError ...
type HTTPError struct {
	code    int
	message string
	error   error
}

// NewHTTPError - constructor
func NewHTTPError(code int, message string, err error) Error {
	return &HTTPError{code, message, err}
}

// NewGenericHTTPError - constructor
func NewGenericHTTPError(err error) Error {
	return &HTTPError{http.StatusInternalServerError, "Oops, something went wrong. Please try again later.", err}
}

// NewUnauthorizedError - constructor
func NewUnauthorizedError(err error) Error {
	return &HTTPError{http.StatusUnauthorized, "Unauthorized request", err}
}

// NewForbiddenError - constructor
func NewForbiddenError(err error, message string) Error {
	return &HTTPError{http.StatusForbidden, message, err}
}

// NewGenericNotFoundError - constructor
func NewGenericNotFoundError() Error {
	return &HTTPError{http.StatusNotFound, "Resource not found", nil}
}

// NewNotFoundError ...
func NewNotFoundError(msg string) Error {
	return &HTTPError{http.StatusNotFound, msg, nil}
}

// NewGenericBadRequestError ...
func NewGenericBadRequestError() Error {
	return &HTTPError{http.StatusBadRequest, "", nil}
}

// NewBadRequestError ...
func NewBadRequestError(msg string) Error {
	return &HTTPError{http.StatusBadRequest, msg, nil}
}

// OriginalError ...
func (e *HTTPError) OriginalError() error {
	return e.error
}

// Code - HttpCode
func (e *HTTPError) Code() int {
	return e.code
}

// Message ...
func (e *HTTPError) Message() string {
	return e.message
}
