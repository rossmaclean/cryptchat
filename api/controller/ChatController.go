package controller

import (
	"cryptchat/model"
	"cryptchat/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ChatsHandler Get all chats for a user
func ChatsHandler(c *gin.Context) {
	var request model.ChatsRequest
	c.BindJSON(&request)
	chats, err := service.GetChats(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	response := model.ChatsResponse{
		Username: request.Username,
		Chats:    chats,
	}
	c.JSON(http.StatusOK, response)
}

// ChatHandler Get all messages for a specific chat
func ChatHandler(c *gin.Context) {
	var request model.ChatRequest
	c.BindJSON(&request)
	chatMessages, err := service.GetChatMessages(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, chatMessages)
}
