package models

type User struct {
	UserId		string `json:"UserId"`
	Name			string `json:"Name"`
	Email			string `json:"Email"`
	Password	string `json:"Password"`
}