package user

import (
	"backend/core/models"
	"backend/utils"
	"context"

	"google.golang.org/api/iterator"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) GetUserByField(field string, value string) (models.User, error, bool) {
	db := utils.GetFirestoreClient()
	ctx := context.Background()

	user := models.User{}
	itr := db.Collection("user").Where(field, "==", value).Documents(ctx)
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

func (repo *UserRepository) DoesUserExistByField(field string, value string) bool {
	db := utils.GetFirestoreClient()
	ctx := context.Background()

	itr := db.Collection("user").Where(field, "==", value).Documents(ctx)
	_, err := itr.Next()

	return err != iterator.Done
}

func (repo *UserRepository) CreateUser(registeringUser models.RegisteringUser) error {
	db := utils.GetFirestoreClient()
	ctx := context.Background()

	_, _, err := db.Collection("user").Add(ctx, registeringUser)

	return err
}
