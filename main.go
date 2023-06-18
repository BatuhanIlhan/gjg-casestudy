package main

import (
	"fmt"
	"github.com/BatuhanIlhan/gjg-casestudy/database"
	"github.com/BatuhanIlhan/gjg-casestudy/handlers"
	"github.com/BatuhanIlhan/gjg-casestudy/repositories"
	"github.com/BatuhanIlhan/gjg-casestudy/services"
	"github.com/BatuhanIlhan/gjg-casestudy/settings"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Initialize database
	db := database.New(database.Config{
		URI:             settings.Cnf.PostgresURI,
		MigrationSource: settings.Cnf.MigrationSource,
		Debug:           settings.Cnf.Debug,
	})

	defer db.Close()
	db.Connect()
	db.Migrate()

	// Repositories
	userRepo := repositories.NewUserRepository(db.DB())
	submissionRepo := repositories.NewSubmissionRepository(db.DB())

	// Services
	userService := services.NewUserService(userRepo)
	submissionService := services.NewSubmissionService(submissionRepo)
	app := fiber.New(fiber.Config{
		IdleTimeout:  settings.Cnf.IdleTimeout,
		ReadTimeout:  settings.Cnf.ReadTimeout,
		WriteTimeout: settings.Cnf.WriteTimeout,
	})

	router := app.Group(settings.Cnf.BaseUrl)

	// Handlers
	handlers.SetupUser(router.Group("/"), handlers.NewUserHandler(userService))
	handlers.SetupSubmission(router.Group("/"), handlers.NewSubmissionHandler(submissionService))

	go func() {
		address := fmt.Sprintf(":%s", settings.Cnf.Port)
		fmt.Println(fmt.Sprintf("Application starting at %v", address))
		if err := app.Listen(address); err != nil {
			log.Fatalf("Application could not started %v", err.Error())
		}
	}()
	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	_ = app.Shutdown()
	fmt.Println("shutdown succeeded")

}
