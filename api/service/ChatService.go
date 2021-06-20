package service

import (
	"cryptchat/model"
)

func GetChats(request model.ChatsRequest) ([]model.Chat, error) {
	chats := make([]model.Chat, 2)
	chat1 := model.Chat{
		ID: "54321",
	}
	chats[0] = chat1

	chat2 := model.Chat{
		ID: "1234",
	}
	chats[1] = chat2
	return chats, nil
}

func GetChatDetails(request model.ChatRequest) (model.ChatDetails, error) {

}
