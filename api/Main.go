package main

import (
	"cryptchat/controller"
	"cryptchat/repository"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Println("Starting Application")
	repository.InitDbSession()
	router := gin.Default()
	//router.LoadHTMLGlob("../frontend/build/*.html")
	//router.Static("/static", "./frontend/build/static")
	//router.GET("/", func(ctx *gin.Context) {
	//	ctx.HTML(http.StatusOK, "index.html", gin.H{})
	//})
	router.POST("/api/v1/signup", controller.SignupHandler)
	router.POST("/api/v1/login", controller.LoginHandler)

	router.POST("/api/v1/chats", controller.GetChatsHandler)
	router.POST("/api/v1/chat_messages", controller.GetChatMessagesHandler)
	router.POST("/api/v1/new_message", controller.SaveChatMessageHandler)

	err := router.Run(":8000")
	if err != nil {
		log.Fatal("Unable to start web server", err)
	}
	log.Println("Application Running")
}
