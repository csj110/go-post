package main

import (
	"blogos/controllers"
	"blogos/database"
	"blogos/middlewares"
	"blogos/models"
	"blogos/repository/crud"
	"fmt"
	"net/http"

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

	r.Use(Cors())
	
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
	rPosts.Use(middlewares.AuthCheck())
	{
		rPosts.GET("", controllers.GetPosts(postRepo))
		rPosts.GET("/:id", controllers.GetPost(postRepo))
		rPosts.POST("", controllers.CreatePost(postRepo))
		rPosts.PUT("/:id", controllers.UpdatePost(postRepo))
		rPosts.DELETE("/:id", controllers.DeletePost(postRepo))
	}

	rAuth := r.Group("/auth")
	{
		rAuth.POST("/login", controllers.Login())
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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
