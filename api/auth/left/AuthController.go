package authleft

import (
	authcore "cryptchat/auth/core"
	"cryptchat/auth/right"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SignupHandler(c *gin.Context) {
	var request authright.SignupRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(400)
		log.Fatal(err)
		return
	}

	err = authcore.Signup(request)
	if err != nil {
		c.Status(500)
		log.Fatal(err)
		return
	}
	c.Status(200)
}

func LoginHandler(c *gin.Context) {
	var request authright.LoginRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(400)
		log.Fatal(err)
		return
	}
	authToken, err := authcore.Login(request)
	if err != nil || authToken == "" {
		c.Status(403)
		return
	}
	c.JSON(http.StatusOK, authToken)
}
