package service

import (
	"cryptchat/model"
	"cryptchat/repository"
)

func GetChats(request model.ChatsRequest) ([]model.Chat, error) {
	chats, err := repository.GetChats(request.Username)
	return chats, err
}

func GetChatMessages(request model.ChatRequest) ([]model.ChatMessage, error) {
	messages, err := repository.GetChatMessages(request.ChatId)
	return messages, err
}
