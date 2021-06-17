package model

type SignupRequest struct {
	Username        string `json:"username"`
	ConfirmUsername string `json:"confirmUsername"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserAuth struct {
	Username       string
	HashedPassword []byte
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
