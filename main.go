package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       string `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func createUser(c *gin.Context) {
	var newUser User
	var err error = c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error parsing request body")
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func main() {
	router := gin.Default()
	router.POST("/register", createUser)

	router.Run("localhost:8080")
}
