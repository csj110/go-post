package repository

import "blogos/models"

type PostRepository interface {
	Save(models.Post) (models.Post, error)
	FindAll() ([]models.Post, error)
	FindById(uint) (models.Post, error)
	UpdatePost(uint, models.Post,uint) (int64, error)
	DeletePost(uint,uint) (int64, error)
}
