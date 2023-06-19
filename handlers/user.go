package handlers

import (
	"github.com/BatuhanIlhan/gjg-casestudy/common"
	"github.com/BatuhanIlhan/gjg-casestudy/common/errors"
	"github.com/BatuhanIlhan/gjg-casestudy/common/transformers"
	"github.com/BatuhanIlhan/gjg-casestudy/models"
	"github.com/BatuhanIlhan/gjg-casestudy/services"
	"github.com/go-openapi/strfmt"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"strconv"
)

const DEFAULTLIMITVALUE = 2
const DEFAULTOFFSETVALUE = 0

type UserHandler struct {
	service                     *services.UserService
	userTransformer             transformers.UserTransformer
	userWithRankTransformer     transformers.UserWithRankTransformer
	userWithRankListTransformer transformers.UserWithRankListTransformer
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service, userTransformer: transformers.User,
		userWithRankTransformer: transformers.UserWithRank, userWithRankListTransformer: transformers.UserWithRankList}
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
	switch err {
	case nil:
		return c.Status(fiber.StatusOK).JSON(h.userTransformer(user))
	//case errors.WalletExist:
	//	return common.RespondCatchedSerivceError(c, err)
	default:
		return common.RespondServiceError(c, err)
	}

}

func (h *UserHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		serviceErr := errors.NewQueryParameterRequired("id")
		return common.RespondServiceError(c, &serviceErr)
	}

	user, err := h.service.Get(c.Context(), id)
	switch err {
	case nil:
		return c.Status(fiber.StatusOK).JSON(h.userWithRankTransformer(user))
	default:
		return common.RespondServiceError(c, err) //500
	}

}

func (h *UserHandler) GetLeaderBoard(c *fiber.Ctx) error {
	limitParam := c.Params("limit")
	var limit int
	if limitParam == "" {
		limit = DEFAULTLIMITVALUE
	} else {
		limit, _ = strconv.Atoi(limitParam)
	}

	offsetParam := c.Params("offset")
	var offset int
	offset, _ = strconv.Atoi(offsetParam)
	if offsetParam == "" {
		offset = DEFAULTOFFSETVALUE
	}

	leaderBoard, err := h.service.GetLeaderBoard(c.UserContext(), limit, offset)
	if err != nil {
		return common.RespondServiceError(c, err)
	}
	total := int64(len(leaderBoard))

	response := models.LeaderBoard{
		PaginatedResponse: models.PaginatedResponse{
			Limit:  int64(limit),
			Offset: int64(offset),
			Total:  total,
		},
		Data: h.userWithRankListTransformer(leaderBoard),
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *UserHandler) GetLeaderBoardByCountry(c *fiber.Ctx) error {
	limitParam := c.Params("limit")
	var limit int
	if limitParam == "" {
		limit = DEFAULTLIMITVALUE
	} else {
		limit, _ = strconv.Atoi(limitParam)
	}

	offsetParam := c.Params("offset")
	var offset int
	offset, _ = strconv.Atoi(offsetParam)
	if offsetParam == "" {
		offset = DEFAULTOFFSETVALUE
	}

	isoCode := c.Params("iso_code")
	if isoCode == "" {
		serviceErr := errors.NewQueryParameterRequired("country_code")
		return common.RespondServiceError(c, &serviceErr)
	}

	leaderBoard, err := h.service.GetLeaderBoard(c.UserContext(), limit, offset, qm.Where("country_code = ?", isoCode))
	if err != nil {
		return common.RespondServiceError(c, err)
	}
	total := int64(len(leaderBoard))

	response := models.LeaderBoard{
		PaginatedResponse: models.PaginatedResponse{
			Limit:  int64(limit),
			Offset: int64(offset),
			Total:  total,
		},
		Data: h.userWithRankListTransformer(leaderBoard),
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
