package util

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt"
)

var (
	ErrJwtMissingToken = errors.New("Missing token")
	ErrJwtInvalidToken = errors.New("Invalid token")
)

func ParseJWTWithClaims(jwtSecret, tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	if tokenString == "" {
		return nil, nil, ErrJwtMissingToken
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, nil, err
	}
	if !token.Valid {
		return nil, nil, ErrJwtInvalidToken
	}

	return token, claims, nil
}
