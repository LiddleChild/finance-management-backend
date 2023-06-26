package databases

import (
	"backend/models"
	"backend/utils"
	"context"

	"google.golang.org/api/iterator"
)

func GetUserByField(field string, value string) (models.User, error, bool) {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	user := models.User{}
	itr := dbClient.Collection("user").Where(field, "==", value).Documents(ctx)
	doc, err := itr.Next()
	if err == iterator.Done {
		return user, nil, false
	} else if err != nil {
		return user, err, false
	}

	doc.DataTo(&user)
	user.UserId = doc.Ref.ID
	return user, nil, true
}

func DoesUserExistByField(field string, value string) bool {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	itr := dbClient.Collection("user").Where(field, "==", value).Documents(ctx)
	_, err := itr.Next()
	
	return err != iterator.Done
}

func CreateUser(registeringUser models.RegisteringUser) error {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	_, _, err := dbClient.Collection("user").Add(ctx, registeringUser)
	
	return err
}