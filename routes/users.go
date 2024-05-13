package routes

import (
	"net/http"
	"udemy/restapi/models"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt parse request data"})
		return
	}

	user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created Successfully"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt save user"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "couldnt authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login succesful"})
}
