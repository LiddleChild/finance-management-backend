package user

import (
	"backend/core/models"
	"backend/core/repository/user"
)

type IUserRepository interface {
	GetUserByField(field string, value string) (models.User, error, bool)
	DoesUserExistByField(field string, value string) bool
	CreateUser(registeringUser models.RegisteringUser) error
}

func NewUserRepository() IUserRepository {
	return user.NewUserRepository()
}
