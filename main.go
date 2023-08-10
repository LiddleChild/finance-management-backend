package main

import (
	"backend/controllers"
	auth_mw "backend/core/middlewares/auth"
	"backend/package/middlewares/auth"
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

	auth := auth.NewMiddleware(auth_mw.NewMock("fmWEwAx6QXokS5xnICKW"))

	app.Get("/", controllers.Ping)

	app.Post("/auth/create_user", controllers.CreateUser)
	app.Post("/auth/login", controllers.Login)
	app.Post("/auth/logout", controllers.Logout)

	app.Get("/transaction/today", auth.RequireAccessToken(controllers.GetTodayTransaction))
	app.Get("/transaction", auth.RequireAccessToken(controllers.GetTransaction))
	app.Post("/transaction", auth.RequireAccessToken(controllers.CreateTransaction))
	app.Patch("/transaction", auth.RequireAccessToken(controllers.UpdateTransaction))
	app.Delete("/transaction", auth.RequireAccessToken(controllers.DeleteTransaction))

	app.Get("/category", auth.RequireAccessToken(controllers.GetCategoryMap))
	app.Post("/category", auth.RequireAccessToken(controllers.CreateCategory))
	app.Patch("/category", auth.RequireAccessToken(controllers.UpdateCategory))
	app.Delete("/category", auth.RequireAccessToken(controllers.DeleteCategory))

	app.Get("/wallet", auth.RequireAccessToken(controllers.GetWalletMap))
	app.Post("/wallet", auth.RequireAccessToken(controllers.CreateWallet))
	app.Patch("/wallet", auth.RequireAccessToken(controllers.UpdateWallet))
	app.Delete("/wallet", auth.RequireAccessToken(controllers.DeleteWallet))

	// Listen to port
	app.Listen(fmt.Sprintf("localhost:%s", utils.GetEnv("PORT", "8080")))

	// Close connection
	utils.CloseFirestoreClient()
}
