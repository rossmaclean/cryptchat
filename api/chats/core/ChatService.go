package chatscore

import (
	chatsright "cryptchat/chats/right"
)

var GetChatsForUser = func(userId string) ([]ChatMongo, error) {
	return chatsright.GetChatsForUserMongo(userId)
}

var GetChatMessages = func(chatId string) (ChatMessageMongo, error) {
	return chatsright.GetChatMessagesMongo(chatId)
}
