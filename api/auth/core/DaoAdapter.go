package authcore

import (
	authright "cryptchat/auth/right"
)

var SaveUser = func(user authright.User) error {
	return authright.SaveUserMongo(user)
}

var FindUser = func(username string) (*authright.User, error) {
	return authright.FindUserMongo(username)
}
