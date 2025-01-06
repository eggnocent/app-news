package entity

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtData struct {
	UserId float64 `json:"user_id"`
	jwt.RegisteredClaims
}