package users

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserId string `json:"userId"`
	Username string `json:"username"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}


type User struct {
	UserId string
	Username string
	Email string
}
