package chatsright

type GetChatsRequestMongo struct {
	UserId string `json:"userId"`
}

type GetChatsResponseMongo struct {
	Chats []Chat `json:"chats"`
}

type GetChatMessagesRequestMongo struct {
	ChatId string `json:"chatId"`
}

type ChatMongo struct {
	ChatId    string   `json:"chatId" bson:"chatId"`
	UserIds   []string `json:"userId" bson:"userIds"`
	Timestamp string   `json:"timestamp" bson:"timestamp"`
}

type GetChatMessagesResponseMongo struct {
	ChatMessages []ChatMessage `json:"chatMessages"`
}

type ChatMessageMongo struct {
	ChatId   string          `json:"chatId" bson:"chatId"`
	Messages []MessagesMongo `json:"messages" bson:"messages"`
}

type MessagesMongo struct {
	UserId    string `json:"userId" bson:"userId"`
	Message   string `json:"message" bson:"message"`
	MessageId string `json:"messageId" bson:"messageId"`
	Timestamp string `json:"timestamp" bson:"timestamp"`
}

type SaveChatMessageRequestMongo struct {
	ChatId  string `json:"chatId"`
	UserId  string `json:"userId"`
	Message string `json:"message"`
}
