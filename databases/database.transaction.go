package databases

import (
	"backend/models"
	"backend/utils"
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func GetTransactionsInRangeByUserId(userId string, startEpoch int64, endEpoch int64) ([]models.Transaction, error) {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

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

func CreateTransaction(userId string, transaction models.Transaction) error {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	_, err := dbClient.Collection("user").
		Doc(userId).
		Collection("transaction").
		Doc(transaction.TransactionId).
		Set(ctx, transaction)

	return err
}

func UpdateTransaction(userId string, transaction models.Transaction) error {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	_, err := dbClient.Collection("user").
		Doc(userId).
		Collection("transaction").
		Doc(transaction.TransactionId).
		Set(ctx, transaction)

	return err
}

func DeleteTransaction(userId string, transaction models.DeletingTransaction) error {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	doc, err := dbClient.Collection("user").
		Doc(userId).
		Collection("transaction").
		Doc(transaction.TransactionId).
		Get(ctx)
	if err != nil {
		return err
	}

	_, err = doc.Ref.Delete(ctx)
	return err
}
