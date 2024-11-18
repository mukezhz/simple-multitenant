package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sachin-gautam/gin-api/helper"
	"github.com/sachin-gautam/gin-api/model"
)

func Register(context *gin.Context) {
	var input model.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
	}
	savedUser, err := user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func Login(contex *gin.Context) {
	var input model.AuthenticationInput

	if err := contex.ShouldBindJSON(&input); err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := model.FinderUserByUsername(input.Username)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = user.ValidatePassword(input.Password)
	if err != nil {
		contex.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	contex.JSON(http.StatusOK, gin.H{"jwt": jwt})

}
