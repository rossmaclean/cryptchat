package repository

import (
	"cryptchat/model"
	"fmt"
	"github.com/gocql/gocql"
)

var GetChats = func(username string) ([]model.Chat, error) {
	var chats []model.Chat
	iter := session.Query(`SELECT user, chat_id, timestamp FROM chats WHERE user=?`,
		username).Consistency(gocql.One).Iter()
	for {
		// New map each iteration
		row := make(map[string]interface{})
		if !iter.MapScan(row) {
			break
		}
		// Do things with row
		chat := model.Chat{
			User:      fmt.Sprintf("%v", row["user"]),
			ChatId:    fmt.Sprintf("%v", row["chat_id"]),
			Timestamp: fmt.Sprintf("%v", row["timestamp"]),
		}
		chats = append(chats, chat)
	}

	var combinedChats []model.Chat
	for x := 0; x < len(chats); x++ {
		chat := new(model.Chat)
		session.Query(`SELECT user, chat_id, timestamp FROM chats WHERE chat_id=?`,
			chats[x].ChatId).Consistency(gocql.One).Scan(&chat.User, &chat.ChatId, &chat.Timestamp)
		combinedChats = append(combinedChats, *chat)
	}

	return combinedChats, nil
}

var GetChatMessages = func(chatId string) ([]model.ChatMessage, error) {
	var chatMessages []model.ChatMessage
	iter := session.Query(`SELECT chat_id, message_id, timestamp, message, user FROM chat_messages WHERE chat_id=?`,
		chatId).Consistency(gocql.One).Iter()
	for {
		// New map each iteration
		row := make(map[string]interface{})
		if !iter.MapScan(row) {
			break
		}
		// Do things with row
		chatMessage := model.ChatMessage{
			ChatId:    fmt.Sprintf("%v", row["chat_id"]),
			MessageId: fmt.Sprintf("%v", row["message_id"]),
			Timestamp: fmt.Sprintf("%v", row["timestamp"]),
			Message:   fmt.Sprintf("%v", row["message"]),
			User:      fmt.Sprintf("%v", row["user"]),
		}
		chatMessages = append(chatMessages, chatMessage)
	}

	return chatMessages, nil
}
