package service

import (
	"cryptchat/model"
	"cryptchat/repository"
	"errors"
	"reflect"
	"testing"
)

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
					return nil, nil
				}

				GenerateFromPassword = func(password []byte, cost int) ([]byte, error) {
					return []byte("hash"), nil
				}

				repository.SaveUserAuth = func(auth model.UserAuth) error {
					return nil
				}
			},
			expectingErr: false,
		},
		{
			name: "Usernames do not match",
			request: model.SignupRequest{
				Username:        "user2",
				ConfirmUsername: "not-matching",
				Password:        "password",
				ConfirmPassword: "password",
			},
			mockFunc: func() {
				repository.GetUserAuth = func(username string) (*model.UserAuth, error) {
					return nil, nil
				}

				GenerateFromPassword = func(password []byte, cost int) ([]byte, error) {
					return []byte("hash"), nil
				}

				repository.SaveUserAuth = func(auth model.UserAuth) error {
					return nil
				}
			},
			expectingErr: true,
		},
		{
			name: "Passwords do not match",
			request: model.SignupRequest{
				Username:        "user3",
				ConfirmUsername: "user3",
				Password:        "password",
				ConfirmPassword: "not-matching",
			},
			mockFunc: func() {
				repository.GetUserAuth = func(username string) (*model.UserAuth, error) {
					return nil, nil
				}

				GenerateFromPassword = func(password []byte, cost int) ([]byte, error) {
					return []byte("hash"), nil
				}

				repository.SaveUserAuth = func(auth model.UserAuth) error {
					return nil
				}
			},
			expectingErr: true,
		},
		{
			name: "User already exists",
			request: model.SignupRequest{
				Username:        "user3",
				ConfirmUsername: "user3",
				Password:        "password",
				ConfirmPassword: "password",
			},
			mockFunc: func() {
				repository.GetUserAuth = func(username string) (*model.UserAuth, error) {
					return &model.UserAuth{
						Username:       "user3",
						HashedPassword: []byte("hash"),
					}, nil
				}

				GenerateFromPassword = func(password []byte, cost int) ([]byte, error) {
					return []byte("hash"), nil
				}

				repository.SaveUserAuth = func(auth model.UserAuth) error {
					return nil
				}
			},
			expectingErr: true,
		},
		{
			name: "Error hashing password",
			request: model.SignupRequest{
				Username:        "user3",
				ConfirmUsername: "user3",
				Password:        "password",
				ConfirmPassword: "password",
			},
			mockFunc: func() {
				repository.GetUserAuth = func(username string) (*model.UserAuth, error) {
					return nil, nil
				}

				GenerateFromPassword = func(password []byte, cost int) ([]byte, error) {
					return nil, errors.New("")
				}

				repository.SaveUserAuth = func(auth model.UserAuth) error {
					return nil
				}
			},
			expectingErr: true,
		},
		{
			name: "Error saving user",
			request: model.SignupRequest{
				Username:        "user3",
				ConfirmUsername: "user3",
				Password:        "password",
				ConfirmPassword: "password",
			},
			mockFunc: func() {
				repository.GetUserAuth = func(username string) (*model.UserAuth, error) {
					return nil, nil
				}

				GenerateFromPassword = func(password []byte, cost int) ([]byte, error) {
					return []byte("hash"), nil
				}

				repository.SaveUserAuth = func(auth model.UserAuth) error {
					return errors.New("")
				}
			},
			expectingErr: true,
		},
	}

	// preserve the original function
	oriGetUserAuth := repository.GetUserAuth
	oriGenerateFromPassword := GenerateFromPassword
	oriSaveUserAuth := repository.SaveUserAuth

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockFunc()
			err := Signup(tc.request)

			errExist := err != nil
			if tc.expectingErr != errExist {
				tt.Errorf("Error expectation not met, wanted [%v], got [%v]", tc.expectingErr, errExist)
			}
		})
	}

	repository.GetUserAuth = oriGetUserAuth
	GenerateFromPassword = oriGenerateFromPassword
	repository.SaveUserAuth = oriSaveUserAuth
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
