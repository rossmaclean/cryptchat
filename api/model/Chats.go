package model

type ChatsRequest struct {
	Username string `json:"username"`
}

type ChatsResponse struct {
	Username string `json:"username"`
	Chats    []Chat `json:"chats"`
}

type ChatRequest struct {
	ChatId string `json:"chatId"`
}

type Chat struct {
	User      string `json:"user"`
	ChatId    string `json:"chatId"`
	Timestamp string `json:"timestamp"`
}

type ChatMessage struct {
	ChatId    string `json:"chatId"`
	MessageId string `json:"messageId"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	User      string `json:"user"`
}
