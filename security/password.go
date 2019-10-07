package security

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (result []byte, err error) {
	return bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword,password string ) (err error){
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
} 
