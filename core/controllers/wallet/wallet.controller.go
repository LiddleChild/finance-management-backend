package wallet

import (
	"backend/core/models"
	"backend/package/repository/wallet"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type WalletController struct {
	repo wallet.IWalletRepository
}

func NewWalletController(repo wallet.IWalletRepository) *WalletController {
	return &WalletController{
		repo,
	}
}

func (con *WalletController) GetWalletMap(c *fiber.Ctx) error {
	userId := c.Locals("UserId").(string)
	walletMap, err := con.repo.GetWalletMapByUserId(userId)
	if err != nil {
		if err != nil {
			fmt.Println(err)
			return c.
				Status(http.StatusConflict).
				SendString(utils.JSONMessage("Couldn't get wallets"))
		}
	}

	jsonStr, err := json.Marshal(walletMap)

	return c.
		Status(http.StatusOK).
		SendString(string(jsonStr))
}

/*
Create wallet
*/
func (con *WalletController) CreateWallet(c *fiber.Ctx) error {
	userId := c.Locals("UserId").(string)

	wallet := models.Wallet{
		WalletId: utils.GenerateUUID(),
	}

	// Parse body
	err := c.BodyParser(&wallet)
	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage("Empty wallet information"))
	}

	// Validate user information
	err = utils.GetValidator().Struct(wallet)
	if err != nil {
		errs := utils.ErrorsToString(utils.TranslateError(err))
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage(
				fmt.Sprintf("Invalid wallet information: %s", strings.Join(errs, ", "))))
	}

	// Create wallet
	err = con.repo.CreateWallet(userId, wallet)
	if err != nil {
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Couldn't create wallet"))
	}

	return c.
		Status(http.StatusCreated).
		SendString(utils.JSONMessage("Wallet created"))
}

/*
Update wallet
*/
func (con *WalletController) UpdateWallet(c *fiber.Ctx) error {
	userId := c.Locals("UserId").(string)

	// Parse body
	wallet := models.Wallet{}
	err := c.BodyParser(&wallet)
	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage("Empty transaction information"))
	}

	// Validate user information
	err = utils.GetValidator().Struct(wallet)
	if err != nil {
		errs := utils.ErrorsToString(utils.TranslateError(err))
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage(
				fmt.Sprintf("Invalid wallet information: %s", strings.Join(errs, ", "))))
	}

	// Update wallet
	err = con.repo.UpdateWallet(userId, wallet)
	if err != nil {
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Couldn't update wallet"))
	}

	return c.
		Status(http.StatusCreated).
		SendString(utils.JSONMessage("Category updated"))
}

/*
Delete wallet
*/
func (con *WalletController) DeleteWallet(c *fiber.Ctx) error {
	userId := c.Locals("UserId").(string)

	// Parse body
	deletingWallet := models.DeletingWallet{}
	err := c.BodyParser(&deletingWallet)
	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage("Empty wallet information"))
	}

	// Validate user information
	err = utils.GetValidator().Struct(deletingWallet)
	if err != nil {
		errs := utils.ErrorsToString(utils.TranslateError(err))
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage(
				fmt.Sprintf("Invalid wallet information: %s", strings.Join(errs, ", "))))
	}

	// Delete wallet
	err = con.repo.DeleteWallet(userId, deletingWallet)
	if err != nil {
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Couldn't delete wallet"))
	}

	return c.
		Status(http.StatusCreated).
		SendString(utils.JSONMessage("Wallet deleted"))
}
