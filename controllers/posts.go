package controllers

import (
	"blogos/models"
	"blogos/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPosts(postRepo repository.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		posts, err := postRepo.FindAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, "")
		} else {
			c.JSON(http.StatusOK, posts)
		}
	}
}

func GetPost(postRepo repository.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		fmt.Printf("%v", id)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"err": "bad request",
			})
			return
		}
		post, err := postRepo.FindById(uint(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"err": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, post)
		}
	}
}

func CreatePost(postRepo repository.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var post models.Post
		json.NewDecoder(c.Request.Body).Decode(&post)
		post, err := postRepo.Save(post)
		if err != nil {
			c.Error(err)
		}
		c.JSON(http.StatusOK, post)
	}
}
