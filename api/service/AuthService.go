package service

import (
	"cryptchat/model"
	"cryptchat/repository"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/thanhpk/randstr"
)

var GenerateFromPassword = bcrypt.GenerateFromPassword

func Signup(request model.SignupRequest) error {
	err := verifyUsernameAndPassword(request)
	if err != nil {
		return err
	}

	uAuth, err := repository.GetUserAuth(request.Username)
	if uAuth != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := GenerateFromPassword([]byte(request.Password), 8)
	if err != nil {
		log.Printf("Unable to hash password for user %s, %o", request.Username, err)
		return errors.New("unable to create account")
	}
	userAuth := model.UserAuth{
		Username:       request.Username,
		HashedPassword: hashedPassword,
	}

	err = repository.SaveUserAuth(userAuth)
	if err != nil {
		log.Printf("Unable to save auth user with username %s, %o", request.Username, err)
		return err
	}
	return nil
}

func verifyUsernameAndPassword(request model.SignupRequest) error {
	if request.Username == request.ConfirmUsername &&
		request.Password == request.ConfirmPassword {
		return nil
	}
	return errors.New("usernames and passwords do not match")
}

var CompareHashAndPassword = bcrypt.CompareHashAndPassword

func Login(request model.LoginRequest) (string, error) {
	authUser, err := repository.GetUserAuth(request.Username)
	if authUser == nil {
		return "", errors.New("user not found")
	}
	err = CompareHashAndPassword(authUser.HashedPassword, []byte(request.Password))
	if err != nil {
		return "", err
	}
	log.Printf("User %s logged in", request.Username)
	return CreateToken(authUser.Username), nil
}

func CreateToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	tokenString, error := token.SignedString([]byte(randstr.String(16)))
	if error != nil {
		fmt.Println(error)
	}
	log.Printf("Token for user %s=%s", username, tokenString)
	return tokenString
}
