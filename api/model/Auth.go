package model

type SignupRequest struct {
	Username        string `json:"username"`
	ConfirmUsername string `json:"confirmUsername"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Email           string `json:"email"`
}

type User struct {
	UserId         string
	Username       string
	Email          string
	HashedPassword []byte
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
