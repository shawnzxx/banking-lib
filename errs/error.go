package errs

import "net/http"

type CustomError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

func (e CustomError) AsMessage() *CustomError {
	return &CustomError{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewAuthError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusForbidden,
		Message: message,
	}
}

func NewDomainError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}

func NewUnexpectedError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
