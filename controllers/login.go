package controllers

import (
	"net/http"
	"blogos/auth"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userDTO UserDTO
		json.NewDecoder(c.Request.Body).Decode(&userDTO)
		toeknString, err := auth.SignIn(userDTO.Username, userDTO.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest,map[string]interface{}{
				"message":err,
			})
		}else{
			c.JSON(http.StatusOK,map[string]string{
				"token":toeknString,
			})
		}
	}
}
