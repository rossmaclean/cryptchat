package chatsright

type GetChatsRequest struct {
	UserId string `json:"userId"`
}

type GetChatsResponse struct {
	Chats []Chat `json:"chats"`
}

type GetChatMessagesRequest struct {
	ChatId string `json:"chatId"`
}

type Chat struct {
	UserId    string `json:"userId"`
	ChatId    string `json:"chatId"`
	MessageId string `json:"messageId"`
	Timestamp string `json:"timestamp"`
}

type GetChatMessagesResponse struct {
	ChatMessages []ChatMessage `json:"chatMessages"`
}

type ChatMessage struct {
	ChatId    string `json:"chatId"`
	UserId    string `json:"userId"`
	Message   string `json:"message"`
	MessageId string `json:"messageId"`
	Timestamp string `json:"timestamp"`
}

type SaveChatMessageRequest struct {
	ChatId  string `json:"chatId"`
	UserId  string `json:"userId"`
	Message string `json:"message"`
}
