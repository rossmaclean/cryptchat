package chatsleft

import (
	chatscore "cryptchat/chats/core"
	"cryptchat/chats/right"
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
	log.Println("Chat handler")
	var request ChatsRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(400)
		log.Fatal(err)
		return
	}

	var chats []chatsright.ChatMongo
	chats, err = chatscore.GetChatsForUser(request.UserId)
	if err != nil {
		c.Status(500)
		log.Fatal(err)
		return
	}
	log.Println("Got chats")
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
	var chatMessages chatsright.ChatMessageMongo
	chatMessages, err = chatscore.GetChatMessages(request.ChatId)
	if err != nil {
		c.Status(500)
		return
	}
	c.JSON(http.StatusOK, chatMessages)
}
