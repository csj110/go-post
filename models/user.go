package models

import (
	"blogos/security"

	"github.com/jinzhu/gorm"
)

type (
	User struct {
		gorm.Model
		Username string `gorm:"unique;not null;index:username" json:"username"`
		Password string `gorm:"size:60;not null;" json:"password"`
		Posts []Post `gorm:"foreignkey:AuthorID" json:"posts,omitempty"`
	}
)

func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
