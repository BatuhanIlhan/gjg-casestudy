package handlers

import "github.com/gofiber/fiber/v2"

func SetupUser(router fiber.Router, handler *UserHandler) {
	router.Post("/user/create", handler.Create)
	router.Get("/user/profile/:id", handler.GetById)
	router.Get("/leaderboard/:limit/:offset", handler.GetLeaderBoard)
	router.Get("/leaderboard", handler.GetLeaderBoard)
	router.Get("/leaderboardByCountry/:iso_code/:limit/:offset", handler.GetLeaderBoardByCountry)
	router.Get("/leaderboardByCountry/:iso_code", handler.GetLeaderBoardByCountry)
}
