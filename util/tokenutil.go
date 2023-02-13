package util

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

const (
	secret = "asdfghjkl"
)

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func CreateToken(userID int64) (string, error) {
	// Created the Claims
	claims := Claims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Created tokens
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signed token and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "converting to string error", err
	}

	return tokenString, nil
}

func DecodeToken(tokenString string) (int64, error) {
	claims := Claims{}

	// Parseing the token
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("token is invalid")
	}

	return claims.UserID, nil
}
