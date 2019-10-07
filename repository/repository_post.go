package repository

import "blogos/models"

type PostRepository interface {
	Save(models.Post) (models.Post, error)
	FindAll() ([]models.Post, error)
	FindById(uint) (models.Post, error)
	// UpdateUser(uint, models.Post) (int64, error)
	// DeleteUser(uint) (int64, error)
}
