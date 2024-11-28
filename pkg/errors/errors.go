package errors

import "net/http"

type AppError struct {
	InternalCode int    `json:"code"`
	StatusCode   int    `json:"status_code"`
	Message      string `json:"message"`
	Details      string `json:"details,omitempty"`
}

var (
	ErrBadRequest     = AppError{InternalCode: 1001, StatusCode: http.StatusBadRequest, Message: "Invalid input"}
	ErrValidation     = AppError{InternalCode: 1002, StatusCode: 422, Message: "Validation error"}
	ErrUnauthorized   = AppError{InternalCode: 1003, StatusCode: 401, Message: "Unauthorized"}
	ErrForbidden      = AppError{InternalCode: 1004, StatusCode: 403, Message: "Forbidden"}
	ErrNotFound       = AppError{InternalCode: 1005, StatusCode: 404, Message: "Resource not found"}
	ErrInternalServer = AppError{InternalCode: 2001, StatusCode: 500, Message: "Internal server error"}
)
