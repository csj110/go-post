package controllers

import (
	"blogos/models"
	"blogos/repository"
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

//GetUsers get user
func GetUsers(userRepo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User
		users, err := userRepo.FindAll()
		if err != nil {
			c.Error(err)
		}
		c.JSON(http.StatusOK, users)
	}
}

func GetUser(userRepo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		id, err := strconv.ParseInt(c.Query("id"), 10, 32)
		user, err = userRepo.FindById(uint(id))
		if err != nil {
			c.Error(err)
		}
		c.JSON(http.StatusOK, user)
	}
}

//CreateUser create user
func CreateUser(userRepo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		json.NewDecoder(c.Request.Body).Decode(&user)
		if user.Username=="" || user.Password== ""{
			c.JSON(http.StatusBadRequest,map[string]string{
				"message":"username and password is not nullable",
			})
			return
		}
		user, err := userRepo.Save(user)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUsers update user
func UpdateUsers(userRepo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var res int64
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		json.NewDecoder(c.Request.Body).Decode(&user)
		res, err = userRepo.UpdateUser(uint(id), user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"updated": res,
			})
		}
	}
}

func DeleteUsers(userRepo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
var user models.User
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		json.NewDecoder(c.Request.Body).Decode(&user)
		_, err = userRepo.DeleteUser(uint(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}else{
			c.JSON(http.StatusNoContent,"")
		}
	}
}
