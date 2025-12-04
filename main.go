package main

import (
	"go-fiber-postgres/controllers"
	"go-fiber-postgres/initializers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	// Public Routes
	app.Post("/signup", controllers.SignUp)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
