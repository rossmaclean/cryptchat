package main

import (
	authleft "cryptchat/auth/left"
	chatsleft "cryptchat/chats/left"
	properties2 "cryptchat/properties"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Application")
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	if properties2.GetEnv() != "local" {
		router.LoadHTMLGlob("../frontend/build/*.html")
		router.Static("/static", "./frontend/build/static")
		router.GET("/", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{})
		})
	}

	router.POST("/api/v1/signup", authleft.SignupHandler)
	router.POST("/api/v1/login", authleft.LoginHandler)

	router.POST("/api/v1/chats", chatsleft.ChatsHandler)
	router.POST("/api/v1/chat_messages", chatsleft.ChatMessageHandler)

	err := router.Run(":8000")
	if err != nil {
		log.Fatal("Unable to start web server", err)
	}
	log.Println("Application Running")
}
