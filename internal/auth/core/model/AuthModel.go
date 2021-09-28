package authcoremodel

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type SignupResponse struct {
	Username string `json:"username"`
	UserId   string `json:"userId"`
	Email    string `json:"email"`
}

type User struct {
	UserId         string
	Username       string
	Email          string
	HashedPassword []byte
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
