package chatscoremodel

type ChatsRequest struct {
	UserId string `json:"userId"`
}

type ChatsResponseM struct {
	Chats []Chat `json:"chats"`
}

type ChatMessagesRequest struct {
	ChatId string `json:"chatId"`
}

// Chat Chats collection
type Chat struct {
	ChatId    string   `json:"chatId" bson:"chatId"`
	UserIds   []string `json:"userId" bson:"userIds"`
	Timestamp string   `json:"timestamp" bson:"timestamp"`
}

type ChatMessagesResponse struct {
	ChatMessages []ChatMessage `json:"chatMessages"`
}

// ChatMessage chatMessage collection
type ChatMessage struct {
	ChatId   string     `json:"chatId" bson:"chatId"`
	Messages []Messages `json:"messages" bson:"messages"`
}

type Messages struct {
	UserId    string `json:"userId" bson:"userId"`
	Message   string `json:"message" bson:"message"`
	MessageId string `json:"messageId" bson:"messageId"`
	Timestamp string `json:"timestamp" bson:"timestamp"`
}

type SaveChatMessageRequest struct {
	ChatId  string `json:"chatId"`
	UserId  string `json:"userId"`
	Message string `json:"message"`
}
