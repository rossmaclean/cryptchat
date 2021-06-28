package controller

import (
	"cryptchat/model"
	"cryptchat/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetChatsHandler(c *gin.Context) {
	var request model.GetChatsRequest
	c.BindJSON(&request)
	chats, err := service.GetChats(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	response := model.GetChatsResponse{
		Chats: chats,
	}
	c.JSON(http.StatusOK, response)
}

func GetChatMessagesHandler(c *gin.Context) {
	var request model.GetChatMessagesRequest
	c.BindJSON(&request)
	chats, err := service.GetChatMessages(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	response := model.GetChatMessagesResponse{
		ChatMessages: chats,
	}
	c.JSON(http.StatusOK, response)
}

func SaveChatMessageHandler(c *gin.Context) {
	var request model.SaveChatMessageRequest
	c.BindJSON(&request)
	err := service.SaveChatMessage(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, "")
}
