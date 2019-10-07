package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Title string `gorm:"size:30;not null;unique;" json:"title"`
	Content string `gorm:"size:255;not null;unique" json:"content"`
	Author User `gorm:"foreignkey:AuthorID" json:"author"`
	AuthorID uint `gorm:"not null" json:"author_id"`
}
