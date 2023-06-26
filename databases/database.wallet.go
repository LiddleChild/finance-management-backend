package databases

import (
	"backend/models"
	"backend/utils"
	"context"

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