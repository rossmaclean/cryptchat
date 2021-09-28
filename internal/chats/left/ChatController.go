package chatsleft

import (
	chatscore "cryptchat/internal/chats/core"
	chatscoremodel "cryptchat/internal/chats/core/model"
	"github.com/gin-gonic/gin"
	"log"
)

type ChatsRequest struct {
	UserId string `json:"userId"`
}

type ChatMessageRequest struct {
	ChatId string `json:"chatId"`
}

func ChatsHandler(g *gin.Context) {
	log.Println("Chat handler")
	var request ChatsRequest
	err := g.BindJSON(&request)
	if err != nil {
		g.Status(400)
		log.Println(err)
		return
	}

	var chats []chatscoremodel.Chat
	chats, err = chatscore.GetChatsForUser(request.UserId)
	if err != nil {
		g.Status(500)
		log.Println(err)
		return
	}
	log.Println("Got chats")
	g.JSON(200, chats)
}

/*func ChatMessageHandler(c *gin.Context) {
	var request ChatMessageRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(400)
		log.Println(err)
		return
	}
	var chatMessages mongo.ChatMessageMongo
	chatMessages, err = chatscore.GetChatMessages(request.ChatId)
	if err != nil {
		c.Status(500)
		return
	}
	c.JSON(http.StatusOK, chatMessages)
}*/
