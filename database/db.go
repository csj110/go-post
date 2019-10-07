package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"

)

func Connect() (db *gorm.DB,err error) {
	return gorm.Open("mysql", "root:19941110@/blog?charset=utf8&parseTime=True&loc=Local")
}
