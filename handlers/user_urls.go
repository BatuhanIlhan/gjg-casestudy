package handlers

import "github.com/gofiber/fiber/v2"

func SetupUser(router fiber.Router, handler *UserHandler) {
	router.Post("/user/create", handler.Create)
}
