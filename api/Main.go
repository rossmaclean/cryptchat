package main

import (
	"cryptchat/controller"
	"cryptchat/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting Application")
	repository.InitDbSession()
	router := gin.Default()
	router.LoadHTMLGlob("../frontend/build/*.html")
	router.Static("/static", "./frontend/build/static")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.POST("/api/v1/signup", controller.SignupHandler)
	router.POST("/api/v1/login", controller.LoginHandler)
	router.Run(":8000")
	log.Println("Application Running")
}
