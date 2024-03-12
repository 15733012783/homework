package server

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

const APP_KEY = "wanan"

func SetToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(APP_KEY))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}

func GetToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if err == nil {
		return token, nil
	} else {
		return nil, err
	}
}
