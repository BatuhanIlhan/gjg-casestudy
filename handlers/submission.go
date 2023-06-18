package handlers

import (
	"github.com/BatuhanIlhan/gjg-casestudy/common"
	"github.com/BatuhanIlhan/gjg-casestudy/common/transformers"
	"github.com/BatuhanIlhan/gjg-casestudy/models"
	"github.com/BatuhanIlhan/gjg-casestudy/services"
	"github.com/go-openapi/strfmt"
	"github.com/gofiber/fiber/v2"
)

type SubmissionHandler struct {
	service               *services.SubmissionService
	submissionTransformer transformers.SubmissionTransformer
}

func NewSubmissionHandler(service *services.SubmissionService) *SubmissionHandler {
	return &SubmissionHandler{service: service, submissionTransformer: transformers.Submission}
}

func (h *SubmissionHandler) Create(c *fiber.Ctx) error {
	request := new(models.CreateSubmissionRequest)
	if err := c.BodyParser(request); err != nil {
		return common.RespondBadRequestPayload(c, err) //400
	}
	if err := request.Validate(strfmt.Default); err != nil {
		return common.RespondValidationError(c, err) //400
	}

	payload := services.SubmissionCreatePayload{
		UserId: request.UserID.String(),
		Score:  *request.ScoreWorth,
	}

	submission, err := h.service.Create(c.Context(), payload)
	switch err {
	case nil:
		return c.Status(fiber.StatusOK).JSON(h.submissionTransformer(submission))
	//case errors.WalletExist:
	//	return common.RespondCatchedSerivceError(c, err)
	default:
		return common.RespondServiceError(c, err)
	}

}
