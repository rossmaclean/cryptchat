package authright

import (
	"cryptchat/internal/auth/core/model"
	"cryptchat/internal/auth/right/mongo"
)

type AuthRepository interface {
	InitDb()
	FindOneByEmail(email string) (authcoremodel.User, error)
	SaveOne(user authcoremodel.User) error
	Ping() error
}

var authRepository AuthRepository

func GetAuthRepository() AuthRepository {
	if authRepository != nil {
		return authRepository
	}
	authRepository = &authrightmongo.MongoAuthRepository{}
	authRepository.InitDb()
	return authRepository
}
