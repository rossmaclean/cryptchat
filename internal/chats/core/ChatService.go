package chatscore

import (
	"cryptchat/internal/chats/core/model"
	"cryptchat/internal/chats/right"
)

func GetChatsForUser(userId string) ([]chatscoremodel.Chat, error) {
	return chatsright.GetChatRepository().FindAllByUserId(userId)
}

//var GetChatMessages = func(chatId string) (mongo.ChatMessageMongo, error) {
//return mongo.GetChatMessagesMongo(chatId)
//}
