package controller

import (
	"cryptchat/model"
	"cryptchat/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignupHandler(c *gin.Context) {
	var request model.SignupRequest
	c.BindJSON(&request)
	service.Signup(request)
	c.JSON(http.StatusOK, "")
}

func LoginHandler(c *gin.Context) {
	var request model.LoginRequest
	c.BindJSON(&request)
	authToken, err := service.Login(request)
	if err != nil {
		c.JSON(http.StatusForbidden, "Error")
	}
	c.JSON(http.StatusOK, authToken)
}
