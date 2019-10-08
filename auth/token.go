package auth

import (
	"blogos/config"
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

// func ExtractToken(tokenString string) string{
	
// }
