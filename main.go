package main

import (
	"github.com/gin-gonic/gin"

	"TeamVapp/AuthServer/auth"
)

func main() {
	router := gin.Default()
	router.POST("/register", auth.CreateUser)

	router.Run("localhost:8080")
}
