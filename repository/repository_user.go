package repository

import "blogos/models"

type UserRepository interface {
	Save(models.User)(models.User,error)
	FindAll()([]models.User,error)
	FindById(uint)(models.User,error)
	UpdateUser(uint,models.User)(int64,error)
	DeleteUser(uint)(int64,error)
}
