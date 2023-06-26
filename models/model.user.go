package models

type User struct {
	UserId		string `json:"UserId"`
	Name			string `json:"Name"     validate:"required"`
	Email			string `json:"Email"    validate:"required"`
	Password	string `json:"Password" validate:"required"`
}

type RegisteringUser struct {
	Name			string `json:"Name"     validate:"required"`
	Email			string `json:"Email"    validate:"required"`
	Password	string `json:"Password" validate:"required"`
}

type UserCredentials struct {
	Email			string `json:"Email"    validate:"required"`
	Password	string `json:"Password" validate:"required"`
}