package errors

import (
	"github.com/gofiber/fiber/v2"
)

const (
	CodeInternalServerError = 100
)

type ServiceError struct {
	Code       int
	Message    string
	HTTPStatus int
}

func New(code int, message string, status int) *ServiceError {
	return &ServiceError{code, message, status}
}

var (
	InternalServerError = New(CodeInternalServerError, "Internal server error", fiber.StatusInternalServerError)
)
