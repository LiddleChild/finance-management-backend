package main

import (
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

	app.Get("/", func (c *fiber.Ctx) error {
		return c.
			Status(200).
			SendString("Hello, World!")
	})

	// Listen to port
	app.Listen(fmt.Sprintf("localhost:%s", utils.GetEnv("PORT", "8080")))
}