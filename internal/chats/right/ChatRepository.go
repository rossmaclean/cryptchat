package chatsright

import (
	"cryptchat/internal/chats/core/model"
	"cryptchat/internal/chats/right/mongo"
)

type ChatRepository interface {
	InitDb()
	FindAllByUserId(userId string) ([]chatscoremodel.Chat, error)
}

var chatRepository ChatRepository

func GetChatRepository() ChatRepository {
	if chatRepository != nil {
		return chatRepository
	}
	chatRepository = &chatsrightmongo.MongoChatRepository{}
	chatRepository.InitDb()
	return chatRepository
}
