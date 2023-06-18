package common

import (
	"github.com/BatuhanIlhan/gjg-casestudy/common/errors"
	"github.com/BatuhanIlhan/gjg-casestudy/models"
	"github.com/gofiber/fiber/v2"
)

const (
	CodeInvalidPayload   int64 = 601
	CodeValidationFailed int64 = 602
)

func RespondBadRequestPayload(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(&models.APIError{Code: CodeInvalidPayload, Error: err.Error()})
}

func RespondValidationError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(&models.APIError{Code: CodeValidationFailed, Error: err.Error()})
}

func RespondServiceError(ctx *fiber.Ctx, err *errors.ServiceError) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(&models.APIError{Code: int64(err.Code), Error: err.Message})
}

func RespondCatchedSerivceError(ctx *fiber.Ctx, err *errors.ServiceError) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(&models.APIError{Code: int64(err.Code), Error: err.Message})
}
