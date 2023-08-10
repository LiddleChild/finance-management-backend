package wallet

import (
	walletCon "backend/core/controllers/wallet"
	"backend/package/repository/wallet"

	"github.com/gofiber/fiber/v2"
)

type IWalletController interface {
	GetWalletMap(c *fiber.Ctx) error
	CreateWallet(c *fiber.Ctx) error
	UpdateWallet(c *fiber.Ctx) error
	DeleteWallet(c *fiber.Ctx) error
}

func NewWalletController(repo wallet.IWalletRepository) IWalletController {
	return walletCon.NewWalletController(repo)
}
