package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
}
