package handlers

import "github.com/gofiber/fiber/v2"

func SetupSubmission(router fiber.Router, handler *SubmissionHandler) {
	router.Post("/score/submit", handler.Create)
}
