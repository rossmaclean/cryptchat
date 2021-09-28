package authleft

import (
	"cryptchat/internal/auth/core"
	"cryptchat/internal/auth/core/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SignupHandler(c *gin.Context) {
	var request authcoremodel.SignupRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(400)
		log.Println(err)
		return
	}

	user, err := authcore.Signup(request)
	if err != nil {
		c.Status(500)
		log.Println(err)
		return
	}

	response, err := authcore.Login(authcoremodel.LoginRequest{
		Email:    user.Email,
		Password: request.Password,
	})

	if err != nil || response.Token == "" {
		c.Status(403)
		return
	}

	c.JSON(200, response)
}

func LoginHandler(c *gin.Context) {
	var request authcoremodel.LoginRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(400)
		log.Println(err)
		return
	}
	response, err := authcore.Login(request)
	if err != nil || response.Token == "" {
		c.Status(403)
		return
	}

	c.JSON(http.StatusOK, response)
}
