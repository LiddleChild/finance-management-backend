package database

import (
	"backend/models"
	"backend/utils"
	"context"

	"google.golang.org/api/iterator"
)

func DoesUserExistByField(field string, value string) bool {
	dbClient := utils.GetFirstoreClient()
	ctx := context.Background()

	itr := dbClient.Collection("user").Where(field, "==", value).Documents(ctx)
	_, err := itr.Next()
	
	return err != iterator.Done
}

func CreateUser(registeringUser models.RegisteringUser) error {
	dbClient := utils.GetFirstoreClient()
	ctx := context.Background()

	_, _, err := dbClient.Collection("user").Add(ctx, registeringUser)
	
	return err
}