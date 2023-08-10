package wallet

import (
	"backend/core/models"
	"backend/core/repository/wallet"
)

type IWalletRepository interface {
	GetWalletMapByUserId(userId string) (map[string]models.Wallet, error)
	DoesWalletExist(userId string, walletId string) bool
	CreateWallet(userId string, wallet models.Wallet) error
	UpdateWallet(userId string, wallet models.Wallet) error
	DeleteWallet(userId string, wallet models.DeletingWallet) error
}

func NewWalletRepository() IWalletRepository {
	return wallet.NewWalletRepository()
}
