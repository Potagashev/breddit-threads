package utils

import (
	"fmt"
	"github.com/Potagashev/breddit/internal/users"
	"github.com/Potagashev/breddit/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(tokenString string, cfg config.Config) (*users.User, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&users.UserClaims{},
		func (token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(cfg.JWTSecret), nil
		},
	)
	if err != nil {
		fmt.Println()
		return nil, err
	}

	if claims, ok := token.Claims.(*users.UserClaims); ok && token.Valid {
		return &users.User{
			UserId: claims.UserId,
			Username: claims.Username,
			Email: claims.Email,
		}, nil
	}
	return nil, fmt.Errorf("invalid token")
}
