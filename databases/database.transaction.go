package databases

import (
	"backend/models"
	"backend/utils"
	"context"

	"google.golang.org/api/iterator"
)

func GetTransactionsByUserId(UserId string) ([]models.Transaction, error) {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	transactionLists := make([]models.Transaction, 0)

	itr := dbClient.Collection("user").Doc(UserId).Collection("transaction").Documents(ctx)
	for {
		transactionDoc, err := itr.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return transactionLists, err
		}

		// Get transaction
		transaction := models.Transaction{}
		transactionDoc.DataTo(&transaction)
		transaction.TransactionId = transactionDoc.Ref.ID

		// Append to transaction lists
		transactionLists = append(transactionLists, transaction)
	}

	return transactionLists, nil
}