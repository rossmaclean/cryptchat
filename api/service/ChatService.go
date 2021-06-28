package service

import (
	"cryptchat/model"
	"cryptchat/repository"
)

func GetChats(request model.GetChatsRequest) ([]model.Chat, error) {
	chats, err := repository.GetChats(request.UserId)
	return chats, err
}

func GetChatMessages(request model.GetChatMessagesRequest) ([]model.ChatMessage, error) {
	messages, err := repository.GetChatMessages(request.ChatId)
	return messages, err
}

func SaveChatMessage(request model.SaveChatMessageRequest) error {
	return repository.SaveChatMessage(request.ChatId, request.UserId, request.Message)
}
