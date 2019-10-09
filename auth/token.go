package auth

import (
	"blogos/config"
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func Createtoken(user_id uint) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    user_id,
		"exp":        time.Now().Add(time.Hour * 6).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.GetKey()))
}

func ValidatToken(tokenString string) (user_id interface{}, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v\n", token.Header["alg"])
		}
		return []byte(config.GetKey()), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["user_id"], nil
	} else {
		return nil, errors.New("token expired/wrong token")
	}
}
