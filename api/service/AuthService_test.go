package service

import (
	"cryptchat/model"
	"cryptchat/repository"
	"errors"
	"reflect"
	"testing"
)

//func Signup(request model.SignupRequest) string {
//	uAuth, err := repository.GetUserAuth(request.Username)
//	if uAuth != nil {
//		return "User already exists"
//	}
//
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 8)
//	if err != nil {
//		log.Fatalf("Unable to hash password for user %s, %o", request.Username, err)
//	}
//	log.Printf("Login request for details: [%s, %s, %s]", request.Username, request.Password, hashedPassword)
//
//	userAuth := new(model.UserAuth)
//	userAuth.Username = request.Username
//	userAuth.HashedPassword = hashedPassword
//	repository.SaveUserAuth(*userAuth)
//	return ""
//}

func Test_Signup(t *testing.T) {
	tests := []struct {
		name         string
		request      model.SignupRequest
		mockFunc     func()
		expectingErr bool
	}{
		{
			name: "All success no error",
			request: model.SignupRequest{
				Username:        "user1",
				ConfirmUsername: "user1",
				Password:        "password",
				ConfirmPassword: "password",
			},
			mockFunc: func() {
				repository.GetUserAuth = func(username string) (*model.UserAuth, error) {
					return &model.UserAuth{
						Username:       "user1",
						HashedPassword: []byte("hashedPassword"),
					}, nil
				}

				CompareHashAndPassword = func(hashedPassword []byte, password []byte) error {
					return nil
				}
			},
			expectedToken: "string",
			expectingErr:  false,
		},
		{
			name: "Password incorrect",
			request: model.LoginRequest{
				Username: "user2",
				Password: "password",
			},
			mockFunc: func() {
				repository.GetUserAuth = func(username string) (*model.UserAuth, error) {
					return &model.UserAuth{
						Username:       "user2",
						HashedPassword: []byte("hashedPassword"),
					}, nil
				}

				CompareHashAndPassword = func(hashedPassword []byte, password []byte) error {
					return errors.New("password incorrect")
				}
			},
			expectedToken: "",
			expectingErr:  true,
		},
		{
			name: "User doesn't exist",
			request: model.LoginRequest{
				Username: "user3",
				Password: "password",
			},
			mockFunc: func() {
				repository.GetUserAuth = func(username string) (*model.UserAuth, error) {
					return nil, nil
				}

				CompareHashAndPassword = func(hashedPassword []byte, password []byte) error {
					return nil
				}
			},
			expectedToken: "",
			expectingErr:  true,
		},
	}

	// preserve the original function
	oriGetUserAuth := repository.GetUserAuth
	oriCompareHashAndPassword := CompareHashAndPassword

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockFunc()
			token, err := Login(tc.request)

			errExist := err != nil
			if tc.expectingErr != errExist {
				tt.Errorf("Error expectation not met, wanted [%v], got [%v]", tc.expectingErr, errExist)
			}

			if !reflect.DeepEqual(tc.expectedToken, token) {
				tt.Errorf("Error, user profile expectation not met, wanted [%+v], got [%+v]", tc.expectedToken, token)
			}
		})
	}

	repository.GetUserAuth = oriGetUserAuth
	CompareHashAndPassword = oriCompareHashAndPassword
}

func Test_Login(t *testing.T) {
	tests := []struct {
		name          string
		request       model.LoginRequest
		mockFunc      func()
		expectedToken string
		expectingErr  bool
	}{
		{
			name: "All success no error",
			request: model.LoginRequest{
				Username: "user1",
				Password: "password",
			},
			mockFunc: func() {
				repository.GetUserAuth = func(username string) (*model.UserAuth, error) {
					return &model.UserAuth{
						Username:       "user1",
						HashedPassword: []byte("hashedPassword"),
					}, nil
				}

				CompareHashAndPassword = func(hashedPassword []byte, password []byte) error {
					return nil
				}
			},
			expectedToken: "string",
			expectingErr:  false,
		},
		{
			name: "Password incorrect",
			request: model.LoginRequest{
				Username: "user2",
				Password: "password",
			},
			mockFunc: func() {
				repository.GetUserAuth = func(username string) (*model.UserAuth, error) {
					return &model.UserAuth{
						Username:       "user2",
						HashedPassword: []byte("hashedPassword"),
					}, nil
				}

				CompareHashAndPassword = func(hashedPassword []byte, password []byte) error {
					return errors.New("password incorrect")
				}
			},
			expectedToken: "",
			expectingErr:  true,
		},
		{
			name: "User doesn't exist",
			request: model.LoginRequest{
				Username: "user3",
				Password: "password",
			},
			mockFunc: func() {
				repository.GetUserAuth = func(username string) (*model.UserAuth, error) {
					return nil, nil
				}

				CompareHashAndPassword = func(hashedPassword []byte, password []byte) error {
					return nil
				}
			},
			expectedToken: "",
			expectingErr:  true,
		},
	}

	// preserve the original function
	oriGetUserAuth := repository.GetUserAuth
	oriCompareHashAndPassword := CompareHashAndPassword

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockFunc()
			token, err := Login(tc.request)

			errExist := err != nil
			if tc.expectingErr != errExist {
				tt.Errorf("Error expectation not met, wanted [%v], got [%v]", tc.expectingErr, errExist)
			}

			if !reflect.DeepEqual(tc.expectedToken, token) {
				tt.Errorf("Error, user profile expectation not met, wanted [%+v], got [%+v]", tc.expectedToken, token)
			}
		})
	}

	repository.GetUserAuth = oriGetUserAuth
	CompareHashAndPassword = oriCompareHashAndPassword
}
