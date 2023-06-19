package errors

import (
	"fmt"
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

func NewQueryParameterRequired(param string) ServiceError {
	return *New(600, fmt.Sprintf("Query parameter required: %v", param), fiber.StatusBadRequest)
}
