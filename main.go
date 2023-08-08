package main

import (
	"backend/controllers"
	"backend/middlewares"
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

	app.Get("/", controllers.Ping)

	app.Post("/auth/create_user", controllers.CreateUser)
	app.Post("/auth/login", controllers.Login)
	app.Post("/auth/logout", controllers.Logout)

	app.Get("/transaction/today", middlewares.RequireAccessToken(controllers.GetTodayTransaction))
	app.Get("/transaction", middlewares.RequireAccessToken(controllers.GetTransaction))
	app.Post("/transaction", middlewares.RequireAccessToken(controllers.CreateTransaction))
	app.Patch("/transaction", middlewares.RequireAccessToken(controllers.UpdateTransaction))
	app.Delete("/transaction", middlewares.RequireAccessToken(controllers.DeleteTransaction))

	app.Get("/category", middlewares.RequireAccessToken(controllers.GetCategoryMap))
	app.Post("/category", middlewares.RequireAccessToken(controllers.CreateCategory))
	app.Patch("/category", middlewares.RequireAccessToken(controllers.UpdateCategory))
	app.Delete("/category", middlewares.RequireAccessToken(controllers.DeleteCategory))

	app.Get("/wallet", middlewares.RequireAccessToken(controllers.GetWalletMap))
	app.Post("/wallet", middlewares.RequireAccessToken(controllers.CreateWallet))
	app.Patch("/wallet", middlewares.RequireAccessToken(controllers.UpdateWallet))
	app.Delete("/wallet", middlewares.RequireAccessToken(controllers.DeleteWallet))

	// Listen to port
	app.Listen(fmt.Sprintf("localhost:%s", utils.GetEnv("PORT", "8080")))

	// Close connection
	utils.CloseFirestoreClient()
}
