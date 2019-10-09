package controllers

import (
	"blogos/models"
	"blogos/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"

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
		if user_id, ok := c.Keys["user_id"].(float64); !ok {
			return
		} else {
			post.AuthorID = uint(user_id)
		}
		post, err := postRepo.Save(post)
		if err != nil {
			c.Error(err)
		}
		c.JSON(http.StatusOK, post)
	}
}

func UpdatePost(postRepo repository.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		if user_id, ok := c.Keys["user_id"].(float64); !ok {
			return
		}
		var post models.Post
		id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
		json.NewDecoder(c.Request.Body).Decode(&post)
		_, err := postRepo.UpdatePost(uint(id), post,user_id)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				c.JSON(http.StatusBadRequest, nil)
			} else {
				c.JSON(http.StatusInternalServerError, nil)
			}
		}
		c.JSON(http.StatusOK, nil)
	}
}
func DeletePost(postRepo repository.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		if user_id, ok := c.Keys["user_id"].(float64); !ok {
			return
		}
		id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
		_, err := postRepo.DeletePost(uint(id),user_id)
		if err != nil {
			c.Error(err)
		}
		c.JSON(http.StatusOK, nil)
	}
}
