package main

import (
	categoryController "backend/core/controllers/category"
	auth "backend/core/middlewares/auth"
	authController "backend/package/controllers/auth"
	statusController "backend/package/controllers/status"
	transactionController "backend/package/controllers/transaction"
	walletController "backend/package/controllers/wallet"
	authMiddleware "backend/package/middlewares/auth"
	categoryRepo "backend/package/repository/category"
	transactionRepo "backend/package/repository/transaction"
	"backend/package/repository/user"
	walletRepo "backend/package/repository/wallet"
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

	// Close connection
	defer utils.CloseFirestoreClient()

	// Create fiber instance
	app := fiber.New()

	auth := authMiddleware.NewMiddleware(auth.NewMock("fmWEwAx6QXokS5xnICKW"))

	userRepo := user.NewUserRepository()
	categoryRepo := categoryRepo.NewCategoryRepository()
	walletRepo := walletRepo.NewWalletRepository()
	transactionRepo := transactionRepo.NewTransactionRepository()

	statusController := statusController.NewStatusController()
	app.Get("/", statusController.Ping)

	authController := authController.NewAuthController(userRepo)
	app.Post("/auth/create_user", authController.CreateUser)
	app.Post("/auth/login", authController.Login)
	app.Post("/auth/logout", authController.Logout)

	transcationController := transactionController.NewTransactionController(transactionRepo, walletRepo, categoryRepo)
	app.Get("/transaction/today", auth.RequireAccessToken(transcationController.GetTodayTransaction))
	app.Get("/transaction", auth.RequireAccessToken(transcationController.GetTransaction))
	app.Post("/transaction", auth.RequireAccessToken(transcationController.CreateTransaction))
	app.Patch("/transaction", auth.RequireAccessToken(transcationController.UpdateTransaction))
	app.Delete("/transaction", auth.RequireAccessToken(transcationController.DeleteTransaction))

	categoryController := categoryController.NewCategoryController(categoryRepo)
	app.Get("/category", auth.RequireAccessToken(categoryController.GetCategoryMap))
	app.Post("/category", auth.RequireAccessToken(categoryController.CreateCategory))
	app.Patch("/category", auth.RequireAccessToken(categoryController.UpdateCategory))
	app.Delete("/category", auth.RequireAccessToken(categoryController.DeleteCategory))

	walletController := walletController.NewWalletController(walletRepo)
	app.Get("/wallet", auth.RequireAccessToken(walletController.GetWalletMap))
	app.Post("/wallet", auth.RequireAccessToken(walletController.CreateWallet))
	app.Patch("/wallet", auth.RequireAccessToken(walletController.UpdateWallet))
	app.Delete("/wallet", auth.RequireAccessToken(walletController.DeleteWallet))

	// Listen to port
	app.Listen(fmt.Sprintf("localhost:%s", utils.GetEnv("PORT", "8080")))
}
