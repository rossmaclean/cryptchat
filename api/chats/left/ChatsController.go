package chatsleft

import (
	chatscore "cryptchat/chats/core"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ChatsRequest struct {
	UserId string `json:"userId"`
}

type ChatMessageRequest struct {
	ChatId string `json:"chatId"`
}

func ChatsHandler(c *gin.Context) {
	var request ChatsRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(400)
		log.Fatal(err)
		return
	}

	var chats []chatscore.ChatMongo
	chats, err = chatscore.GetChatsForUser(request.UserId)
	if err != nil {
		c.Status(500)
		log.Fatal(err)
		return
	}
	c.JSON(200, chats)
}

func ChatMessageHandler(c *gin.Context) {
	var request ChatMessageRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(400)
		log.Fatal(err)
		return
	}
	var chatMessages chatscore.ChatMessageMongo
	chatMessages, err = chatscore.GetChatMessages(request.ChatId)
	if err != nil {
		c.Status(500)
		return
	}
	c.JSON(http.StatusOK, chatMessages)
}
