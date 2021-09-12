package chatscore

import (
	chatsright "cryptchat/chats/right"
)

var GetChatsForUser = func(userId string) ([]chatsright.ChatMongo, error) {
	return chatsright.GetChatsForUserMongo(userId)
}

var GetChatMessages = func(chatId string) (chatsright.ChatMessageMongo, error) {
	return chatsright.GetChatMessagesMongo(chatId)
}
