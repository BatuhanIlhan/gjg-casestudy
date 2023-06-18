package handlers

import (
	"fmt"
	"github.com/BatuhanIlhan/gjg-casestudy/common"
	"github.com/BatuhanIlhan/gjg-casestudy/common/transformers"
	"github.com/BatuhanIlhan/gjg-casestudy/models"
	"github.com/BatuhanIlhan/gjg-casestudy/services"
	"github.com/go-openapi/strfmt"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service         *services.UserService
	userTransformer transformers.UserTransformer
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service, userTransformer: transformers.User}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	request := new(models.CreateUserRequest)
	if err := c.BodyParser(request); err != nil {
		return common.RespondBadRequestPayload(c, err) //400
	}
	if err := request.Validate(strfmt.Default); err != nil {
		return common.RespondValidationError(c, err) //400
	}

	payload := services.UserCreatePayload{
		DisplayName: *request.DisplayName,
		CountryCode: request.CountryCode,
		Points:      request.Points,
	}

	user, err := h.service.Create(c.Context(), payload)
	fmt.Println(user)
	switch err {
	case nil:
		return c.Status(fiber.StatusOK).JSON(h.userTransformer(user))
	//case errors.WalletExist:
	//	return common.RespondCatchedSerivceError(c, err)
	default:
		return common.RespondServiceError(c, err)
	}

}
