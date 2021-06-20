package model

type ChatsRequest struct {
	Username string `json:"username"`
}

type ChatsResponse struct {
	Username string `json:"username"`
	Chats    []Chat `json:"chats"`
}

type ChatRequest struct {
	ID string `json:"id"`
}

type Chat struct {
	ID string `json:"id"`
}
