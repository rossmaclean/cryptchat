package main

import (
	"cryptchat/internal"
	authcore "cryptchat/internal/auth/core"
	authleft "cryptchat/internal/auth/left"
	chatsleft "cryptchat/internal/chats/left"
	healthleft "cryptchat/internal/health/left"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Application")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	if internal.GetEnv() != "local" {
		router.LoadHTMLGlob("/app/code/frontend/build/*.html")
		router.Static("/static", "/app/code/frontend/build/static")
		router.GET("/", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{})
		})
	}

	router.POST("/api/v1/signup", authleft.SignupHandler)
	router.POST("/api/v1/login", authleft.LoginHandler)

	router.GET("/api/v1/health", healthleft.HealthHandler)

	router.POST("/api/v1/chats", authcore.IsAuth(), chatsleft.ChatsHandler)
	//router.POST("/api/v1/chat_messages", authcore.IsAuth(), chatsleft.ChatMessageHandler)

	err := router.Run(":8000")
	if err != nil {
		log.Fatal("Unable to start web server", err)
	}
	log.Println("Application Running")
}
