package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       string `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var newUser User
	var err error = c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error parsing request body")
		return
	}
	c.JSON(http.StatusCreated, newUser)
}
