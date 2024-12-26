package users

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserClaims struct {
	UserId uuid.UUID `json:"userId"`
	Username string `json:"username"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}


type User struct {
	UserId uuid.UUID
	Username string
	Email string
}
