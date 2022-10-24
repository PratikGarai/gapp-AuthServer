package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"GApp/AuthServer/auth"
	"GApp/AuthServer/db"
	"GApp/AuthServer/middlewares"
	"GApp/AuthServer/test"
)

func main() {

	operateMode := flag.String("mode", "deploy", "mode of operation. Options : ['deploy', 'test']")
	flag.Parse()
	if *operateMode == "deploy" {
		fmt.Println("Deploy Mode")
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

		err = db.ConnectDb()
		if err != nil {
			fmt.Println("Can't connect to DB")
			return
		} else {
			fmt.Println("Connected to DB")
		}
		router := gin.Default()
		router.SetTrustedProxies(nil)
		router.Use(middlewares.SetJsonTypeMiddleware())
		router.POST("/register", auth.CreateUser)

		fmt.Printf("Connecting to %s at PORT %d\n", SERVER_BIND_ADDRESS, SERVER_BIND_PORT)
		router.Run(fmt.Sprintf("%s:%d", SERVER_BIND_ADDRESS, SERVER_BIND_PORT))

	} else if *operateMode == "test" {
		fmt.Println("Test Mode")
		err := godotenv.Load(".env.test")
		if err != nil {
			fmt.Println("Error loading .env file")
			return
		}
		test.TestMain()
	} else {
		fmt.Println("Invalid Mode")
	}

}
