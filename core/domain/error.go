package domain

import (
	"net/http"
)

type AppError struct {
	Message string
	Code    int
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	ErrNotFound = &AppError{
		Message: "Resource not found",
		Code:    http.StatusNotFound,
	}
	ErrInternalServerError = &AppError{
		Message: "Internal server error",
		Code:    http.StatusInternalServerError,
	}
	ErrDatabaseTimeOut = &AppError{
		Message: "Database operation timed out",
		Code:    http.StatusGatewayTimeout,
	}
	ErrUserAlreadyExists = &AppError{
		Message: "User already exists",
		Code:    http.StatusConflict,
	}
)
