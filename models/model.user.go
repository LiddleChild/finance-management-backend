package models

type User struct {
	UserId		string `json:"UserId"`
	Name			string `json:"Name"`
	Email			string `json:"Email"`
	Password	string `json:"Password"`
}

type RegisteringUser struct {
	Name			string `json:"Name"     validate:"required"`
	Email			string `json:"Email"    validate:"required,email"`
	Password	string `json:"Password" validate:"required,min=8"`
}

type UserCredentials struct {
	Email			string `json:"Email"    validate:"required,email"`
	Password	string `json:"Password" validate:"required,min=8"`
}