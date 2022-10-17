package main

import (
	"github.com/gin-gonic/gin"

	"TeamVapp/AuthServer/auth"
	"TeamVapp/AuthServer/middlewares"
)

func main() {
	router := gin.Default()
	router.Use(middlewares.SetJsonTypeMiddleware())
	router.POST("/register", auth.CreateUser)

	router.Run("localhost:8080")
}
