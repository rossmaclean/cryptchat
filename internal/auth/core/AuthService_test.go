package authcore

//
//import (
//	"errors"
//	"reflect"
//	"testing"
//)
//
//func Test_Signup(t *testing.T) {
//	tests := []struct {
//		name         string
//		request      authright.SignupRequest
//		mockFunc     func()
//		expectingErr bool
//	}{
//		{
//			name: "All success no error",
//			request: authright.SignupRequest{
//				Username:        "user1",
//				ConfirmUsername: "user1",
//				Password:        "password",
//				ConfirmPassword: "password",
//				Email:           "user@mail.com",
//			},
//			mockFunc: func() {
//				FindUser = func(username string) (*authright.User, error) {
//					return nil, nil
//				}
//
//				GenerateFromPassword = func(password []byte, cost int) ([]byte, error) {
//					return []byte("hash"), nil
//				}
//
//				SaveUser = func(auth authright.User) error {
//					return nil
//				}
//			},
//			expectingErr: false,
//		},
//		{
//			name: "Usernames do not match",
//			request: authright.SignupRequest{
//				Username:        "user2",
//				ConfirmUsername: "not-matching",
//				Password:        "password",
//				ConfirmPassword: "password",
//				Email:           "user2@mail.com",
//			},
//			mockFunc: func() {
//				FindUser = func(username string) (*authright.User, error) {
//					return nil, nil
//				}
//
//				GenerateFromPassword = func(password []byte, cost int) ([]byte, error) {
//					return []byte("hash"), nil
//				}
//
//				SaveUser = func(auth authright.User) error {
//					return nil
//				}
//			},
//			expectingErr: true,
//		},
//		{
//			name: "Passwords do not match",
//			request: authright.SignupRequest{
//				Username:        "user3",
//				ConfirmUsername: "user3",
//				Password:        "password",
//				ConfirmPassword: "not-matching",
//				Email:           "user3@mail.com",
//			},
//			mockFunc: func() {
//				FindUser = func(username string) (*authright.User, error) {
//					return nil, nil
//				}
//
//				GenerateFromPassword = func(password []byte, cost int) ([]byte, error) {
//					return []byte("hash"), nil
//				}
//
//				SaveUser = func(auth authright.User) error {
//					return nil
//				}
//			},
//			expectingErr: true,
//		},
//		{
//			name: "User already exists",
//			request: authright.SignupRequest{
//				Username:        "user3",
//				ConfirmUsername: "user3",
//				Password:        "password",
//				ConfirmPassword: "password",
//				Email:           "user3@mail.com",
//			},
//			mockFunc: func() {
//				FindUser = func(username string) (*authright.User, error) {
//					return &authright.User{
//						UserId:         "bhd32",
//						Username:       "user3",
//						HashedPassword: []byte("hash"),
//						Email:          "user3@mail.com",
//					}, nil
//				}
//
//				GenerateFromPassword = func(password []byte, cost int) ([]byte, error) {
//					return []byte("hash"), nil
//				}
//
//				SaveUser = func(auth authright.User) error {
//					return nil
//				}
//			},
//			expectingErr: true,
//		},
//		{
//			name: "Error hashing password",
//			request: authright.SignupRequest{
//				Username:        "user3",
//				ConfirmUsername: "user3",
//				Password:        "password",
//				ConfirmPassword: "password",
//				Email:           "user3@mail.com",
//			},
//			mockFunc: func() {
//				FindUser = func(username string) (*authright.User, error) {
//					return nil, nil
//				}
//
//				GenerateFromPassword = func(password []byte, cost int) ([]byte, error) {
//					return nil, errors.New("")
//				}
//
//				SaveUser = func(auth authright.User) error {
//					return nil
//				}
//			},
//			expectingErr: true,
//		},
//		{
//			name: "Error saving user",
//			request: authright.SignupRequest{
//				Username:        "user3",
//				ConfirmUsername: "user3",
//				Password:        "password",
//				ConfirmPassword: "password",
//				Email:           "user3@mail.com",
//			},
//			mockFunc: func() {
//				FindUser = func(username string) (*authright.User, error) {
//					return nil, nil
//				}
//
//				GenerateFromPassword = func(password []byte, cost int) ([]byte, error) {
//					return []byte("hash"), nil
//				}
//
//				SaveUser = func(auth authright.User) error {
//					return errors.New("")
//				}
//			},
//			expectingErr: true,
//		},
//	}
//
//	// preserve the original function
//	oriGetUser := FindUser
//	oriGenerateFromPassword := GenerateFromPassword
//	oriSaveUser := SaveUser
//
//	for _, tc := range tests {
//		t.Run(tc.name, func(tt *testing.T) {
//			tc.mockFunc()
//			err := Signup(tc.request)
//
//			errExist := err != nil
//			if tc.expectingErr != errExist {
//				tt.Errorf("Error expectation not met, wanted [%v], got [%v]", tc.expectingErr, errExist)
//			}
//		})
//	}
//
//	FindUser = oriGetUser
//	GenerateFromPassword = oriGenerateFromPassword
//	SaveUser = oriSaveUser
//}
//
//func Test_Login(t *testing.T) {
//	tests := []struct {
//		name          string
//		request       authright.LoginRequest
//		mockFunc      func()
//		expectedToken string
//		expectingErr  bool
//	}{
//		//{
//		//	name: "All success no error",
//		//	request: model.LoginRequest{
//		//		Username: "user1",
//		//		Password: "password",
//		//	},
//		//	mockFunc: func() {
//		//		repository.GetUser = func(username string) (*model.User, error) {
//		//			return &model.User{
//		//				UserId:         "abc123",
//		//				Username:       "user1",
//		//				HashedPassword: []byte("hashedPassword"),
//		//				Email:          "user1@mail.com",
//		//			}, nil
//		//		}
//		//
//		//		CompareHashAndPassword = func(hashedPassword []byte, password []byte) error {
//		//			return nil
//		//		}
//		//	},
//		//	expectedToken: "string",
//		//	expectingErr:  false,
//		//},
//		{
//			name: "Password incorrect",
//			request: authright.LoginRequest{
//				Username: "user2",
//				Password: "password",
//			},
//			mockFunc: func() {
//				FindUser = func(username string) (*authright.User, error) {
//					return &authright.User{
//						UserId:         "ads23",
//						Username:       "user2",
//						HashedPassword: []byte("hashedPassword"),
//						Email:          "user2@mail.com",
//					}, nil
//				}
//
//				CompareHashAndPassword = func(hashedPassword []byte, password []byte) error {
//					return errors.New("password incorrect")
//				}
//			},
//			expectedToken: "",
//			expectingErr:  true,
//		},
//		{
//			name: "User doesn't exist",
//			request: authright.LoginRequest{
//				Username: "user3",
//				Password: "password",
//			},
//			mockFunc: func() {
//				FindUser = func(username string) (*authright.User, error) {
//					return nil, nil
//				}
//
//				CompareHashAndPassword = func(hashedPassword []byte, password []byte) error {
//					return nil
//				}
//			},
//			expectedToken: "",
//			expectingErr:  true,
//		},
//	}
//
//	// preserve the original function
//	oriGetUser := FindUser
//	oriCompareHashAndPassword := CompareHashAndPassword
//
//	for _, tc := range tests {
//		t.Run(tc.name, func(tt *testing.T) {
//			tc.mockFunc()
//			token, err := Login(tc.request)
//
//			errExist := err != nil
//			if tc.expectingErr != errExist {
//				tt.Errorf("Error expectation not met, wanted [%v], got [%v]", tc.expectingErr, errExist)
//			}
//
//			if !reflect.DeepEqual(tc.expectedToken, token) {
//				tt.Errorf("Error, user profile expectation not met, wanted [%+v], got [%+v]", tc.expectedToken, token)
//			}
//		})
//	}
//
//	FindUser = oriGetUser
//	CompareHashAndPassword = oriCompareHashAndPassword
//}
