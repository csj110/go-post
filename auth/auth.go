package auth

import (
	"blogos/database"
	"blogos/models"
	"blogos/security"
	"errors"
)

func SignIn(username string, password string) (string, error) {
	db, err := database.Connect()
	if err != nil {
		return "canot connect to database", err
	}
	defer db.Close()
	
	var user models.User
	if err := db.Debug().Model(models.User{}).Where("username = ?", username).Take(&user).Error; err != nil {
		return "", err
	}

	if err = security.VerifyPassword(user.Password, password); err != nil {
		return "", errors.New("wrong pass")
	}

	return Createtoken(user.ID)
}
