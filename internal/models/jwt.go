package models

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Admin    bool   `json:"user"`
	jwt.RegisteredClaims
}
