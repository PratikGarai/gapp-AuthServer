package auth

import (
	"GApp/AuthServer/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserOutput struct {
	Id       string `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var userInput UserInput
	var err = c.BindJSON(&userInput)
	if err != nil {
		utils.HandleJsonParsingError(c, err)
		return
	}
	c.JSON(http.StatusCreated, userInput)
}
