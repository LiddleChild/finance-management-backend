package models

import "github.com/golang-jwt/jwt/v5"

type JWTClaim struct {
	UserId string `json:"UserId"`
	jwt.RegisteredClaims
}