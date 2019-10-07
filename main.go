package main

import (
	"blogos/controllers"
	"blogos/database"
	"blogos/models"
	"blogos/repository/crud"
	"fmt"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	start()
}

func start() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("can not conenct to database")
	}
	defer db.Close()
	r := gin.Default()
	userRepo := crud.NewRepositoryUserCRUD(db)
	postRepo := crud.NewRepositoryPostCRUD(db)
	rUsers := r.Group("/users")
	{
		rUsers.GET("", controllers.GetUsers(userRepo))
		rUsers.GET("/:id", controllers.GetUser(userRepo))
		rUsers.POST("", controllers.CreateUser(userRepo))
		rUsers.PUT("/:id", controllers.UpdateUsers(userRepo))
		rUsers.DELETE("/:id", controllers.DeleteUsers(userRepo))
	}

	rPosts := r.Group("/posts")
	{
		rPosts.GET("", controllers.GetPosts(postRepo))
		rPosts.GET("/:id", controllers.GetPost(postRepo))
		rPosts.POST("", controllers.CreatePost(postRepo))
	}

	r.Run()
}

func migrate() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("can not conenct to database")
	}
	defer db.Close()

	if err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error; err != nil {
		log.Fatal("can not drop database")
	}
	if err := db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error; err != nil {
		log.Fatal("can not auto migrate ")
	}

	if err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error; err != nil {
		log.Fatal("canot add foreign key")
	}

	user := models.User{Username: "Cai Shengjian", Password: "12345678"}
	err = db.Debug().Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatal("can not add user")
	}
	post := models.Post{Title: "title", Content: "Hello world", Author: user}
	if err := db.Debug().Model(&models.Post{}).Create(&post).Error; err != nil {
		log.Fatal("can not add post")
	}
	db.Debug().Model(&user).Related(&user.Posts, "author_id")
	fmt.Println(user)
}
