package routes

import (
	"net/http"

	"example.com/restapi/models"
	"example.com/restapi/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messgage": "Could not parse request data."})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messgage": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created succesfully"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messgage": "User could not be retrieved."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Login unsuccessful"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Login unsuccessful"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Log in successful", "token": token})
}
