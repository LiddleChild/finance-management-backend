package databases

import (
	"backend/models"
	"backend/utils"
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func GetWalletMapByUserId(userId string) (map[string]models.Wallet, error) {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	walletMap := map[string]models.Wallet{}

	itr := dbClient.Collection("user").Doc(userId).Collection("wallet").Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return walletMap, err
		}

		wallet := models.Wallet{}
		doc.DataTo(&wallet)

		// Set id
		wallet.WalletId = doc.Ref.ID

		walletMap[wallet.WalletId] = wallet
	}

	return walletMap, nil
}

func DoesWalletExist(userId string, walletId string) bool {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	_, err := dbClient.Collection("user").Doc(userId).Collection("wallet").Doc(walletId).Get(ctx)

	return err == nil
}

func CreateWallet(userId string, wallet models.Wallet) error {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	_, err := dbClient.Collection("user").
		Doc(userId).
		Collection("wallet").
		Doc(wallet.WalletId).
		Set(ctx, wallet)

	return err
}

func UpdateWallet(userId string, wallet models.Wallet) error {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	_, err := dbClient.Collection("user").
		Doc(userId).
		Collection("wallet").
		Doc(wallet.WalletId).
		Update(ctx, []firestore.Update{
			{
				Path:  "Label",
				Value: wallet.Label,
			}, {
				Path:  "Color",
				Value: wallet.Color,
			},
		})

	return err
}

func DeleteWallet(userId string, wallet models.DeletingWallet) error {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	doc, err := dbClient.Collection("user").
		Doc(userId).
		Collection("wallet").
		Doc(wallet.WalletId).
		Get(ctx)
	if err != nil {
		return err
	}

	_, err = doc.Ref.Delete(ctx)
	return err
}
