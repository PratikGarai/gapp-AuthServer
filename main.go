package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"TeamVapp/AuthServer/auth"
	"TeamVapp/AuthServer/middlewares"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	var SERVER_BIND_PORT int
	var SERVER_BIND_ADDRESS string

	SERVER_BIND_ADDRESS = os.Getenv("SERVER_BIND_URL")
	if SERVER_BIND_ADDRESS == "" {
		SERVER_BIND_ADDRESS = "localhost"
	}
	SERVER_BIND_PORT, err = strconv.Atoi(os.Getenv("SERVER_BIND_PORT"))
	if err != nil {
		fmt.Println("Error parsing SERVER_BIND_PORT to integer")
		return
	}

	router := gin.Default()
	router.Use(middlewares.SetJsonTypeMiddleware())
	router.POST("/register", auth.CreateUser)

	fmt.Printf("Connecting to %s at PORT %d", SERVER_BIND_ADDRESS, SERVER_BIND_PORT)
	router.Run(fmt.Sprintf("%s:%d", SERVER_BIND_ADDRESS, SERVER_BIND_PORT))
}
