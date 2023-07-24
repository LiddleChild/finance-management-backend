package databases

import (
	"backend/models"
	"backend/utils"
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func GetTransactionsByUserId(userId string, month int, year int, month_range int) ([]models.Transaction, error) {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	if month < 0 || year < 0 || month_range < 0 {
		return GetAllTransactionsByUserId(userId)
	}
	
	startEpoch, endEpoch := utils.GetMonthsRange(month, year, month_range)

	transactionLists := make([]models.Transaction, 0)

	itr := dbClient.Collection("user").
		Doc(userId).
		Collection("transaction").
		Where("Timestamp", ">=", startEpoch).
		Where("Timestamp", "<=", endEpoch).
		OrderBy("Timestamp", firestore.Desc).
		Documents(ctx)
	
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

func GetAllTransactionsByUserId(userId string) ([]models.Transaction, error) {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	transactionLists := make([]models.Transaction, 0)

	itr := dbClient.Collection("user").
		Doc(userId).
		Collection("transaction").
		OrderBy("Timestamp", firestore.Desc).
		Documents(ctx)
	
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

func CreateTransaction(userId string, creatingTransaction models.CreatingTransaction) error {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	_, _, err := dbClient.Collection("user").Doc(userId).Collection("transaction").Add(ctx, creatingTransaction)

	return err
}