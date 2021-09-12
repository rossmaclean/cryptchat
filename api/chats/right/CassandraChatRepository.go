package chatsright

import (
	"cryptchat/chats/core"
)

var GetChats = func(userId string) ([]chatscore.Chat, error) {
	//var chats []model.Chat
	//iter := right.session.Query(`SELECT user_id, chat_id, message_id, timestamp FROM chats WHERE user_id=?`,
	//	userId).Consistency(gocql.One).Iter()
	//for {
	//	// New map each iteration
	//	row := make(map[string]interface{})
	//	if !iter.MapScan(row) {
	//		break
	//	}
	//	// Do things with row
	//	chat := model.Chat{
	//		UserId:    fmt.Sprintf("%v", row["user_id"]),
	//		ChatId:    fmt.Sprintf("%v", row["chat_id"]),
	//		MessageId: fmt.Sprintf("%v", row["message_id"]),
	//		Timestamp: fmt.Sprintf("%v", row["timestamp"]),
	//	}
	//	chats = append(chats, chat)
	//}
	//
	//return chats, nil

	return nil, nil
}

var GetChatMessages = func(chatId string) ([]chatscore.ChatMessage, error) {
	//var chatMessages []model.ChatMessage
	////(chat_id), user_id, message_id, timestamp)
	//iter := right.session.Query(`SELECT chat_id, user_id, message_id, message, timestamp FROM chat_messages WHERE chat_id=?`,
	//	chatId).Consistency(gocql.One).Iter()
	//for {
	//	// New map each iteration
	//	row := make(map[string]interface{})
	//	if !iter.MapScan(row) {
	//		break
	//	}
	//	// Do things with row
	//	chatMessage := model.ChatMessage{
	//		ChatId:    fmt.Sprintf("%v", row["chat_id"]),
	//		UserId:    fmt.Sprintf("%v", row["user_id"]),
	//		MessageId: fmt.Sprintf("%v", row["message_id"]),
	//		Message:   fmt.Sprintf("%v", row["message"]),
	//		Timestamp: fmt.Sprintf("%v", row["timestamp"]),
	//	}
	//	chatMessages = append(chatMessages, chatMessage)
	//}
	//
	//return chatMessages, nil

	return nil, nil
}

var SaveChatMessage = func(chatId string, userId string, message string) error {
	//messageId, _ := gocql.RandomUUID()
	//timestamp := time.Now().String()
	//if err := right.session.Query(`INSERT INTO chat_messages (message_id, chat_id, user_id, message, timestamp) VALUES (?, ?, ?, ?, ?)`,
	//	messageId.String(), chatId, userId, message, timestamp).Exec(); err != nil {
	//	return err
	//}
	//return nil

	return nil
}
