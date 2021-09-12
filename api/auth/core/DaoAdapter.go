package authcore

import (
	authright "cryptchat/auth/right"
)

var SaveUser = func(user User) error {
	return authright.SaveUserMongo(user)
	//return authright.SaveUserCassandra(user)
}

var FindUser = func(username string) (*User, error) {
	return authright.FindUserMongo(username)
	//return authright.FindUserCassandra(username)
}
