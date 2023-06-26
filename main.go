package main

import (
	"backend/controllers"
	"backend/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load dotenv
	godotenv.Load()

	// Create fiber instance
	app := fiber.New()

	app.Post("/auth/create_user", controllers.CreateUser)

	// Listen to port
	app.Listen(fmt.Sprintf("localhost:%s", utils.GetEnv("PORT", "8080")))
}