package models

type Wallet struct {
	WalletId string `json:"WalletId"`
	Color    int64  `json:"Color"`
	Label    string `json:"Label"`
}

type DeletingWallet struct {
	WalletId string `json:"WalletId" validate:"required"`
}
