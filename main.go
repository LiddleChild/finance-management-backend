package main

import (
	"backend/controllers"
	"backend/utils"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Initiate Firestore connect
	err = utils.InitiateFirestoreClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create fiber instance
	app := fiber.New()

	app.Post("/auth/create_user", controllers.CreateUser)
	app.Post("/auth/login", controllers.Login)
	app.Post("/auth/logout", controllers.Logout)

	// Listen to port
	app.Listen(fmt.Sprintf("localhost:%s", utils.GetEnv("PORT", "8080")))

	// Close connection
	utils.CloseFirestoreClient()
}