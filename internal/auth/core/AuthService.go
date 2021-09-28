package authcore

import (
	"cryptchat/internal/auth/core/model"
	"cryptchat/internal/auth/right"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var GenerateFromPassword = bcrypt.GenerateFromPassword
var CompareHashAndPassword = bcrypt.CompareHashAndPassword
var audience = "cryptchatusers"
var issuer = "cryptchat.rossmac.co.uk"
var mySigningKey = []byte(os.Getenv("SECRET_KEY"))

func Signup(request authcoremodel.SignupRequest) (authcoremodel.SignupResponse, error) {
	user, err := authright.GetAuthRepository().FindOneByEmail(request.Email)
	if user.UserId != "" {
		return authcoremodel.SignupResponse{}, errors.New("user already exists")
	}

	hashedPassword, err := GenerateFromPassword([]byte(request.Password), 8)
	if err != nil {
		log.Printf("Unable to hash password for user %s, %o", request.Username, err)
		return authcoremodel.SignupResponse{}, errors.New("unable to create account")
	}
	userAuth := authcoremodel.User{
		UserId:         uuid.New().String(),
		Username:       request.Username,
		HashedPassword: hashedPassword,
		Email:          request.Email,
	}

	err = authright.GetAuthRepository().SaveOne(userAuth)
	if err != nil {
		log.Printf("Unable to save auth user with username %s, %o", request.Username, err)
		return authcoremodel.SignupResponse{}, err
	}
	return authcoremodel.SignupResponse{
		Username: request.Username,
		UserId:   userAuth.UserId,
		Email:    request.Email,
	}, err
}

func Login(request authcoremodel.LoginRequest) (authcoremodel.LoginResponse, error) {
	authUser, err := authright.GetAuthRepository().FindOneByEmail(request.Email)
	if authUser.UserId == "" {
		return authcoremodel.LoginResponse{}, errors.New("user not found")
	}
	err = CompareHashAndPassword(authUser.HashedPassword, []byte(request.Password))
	if err != nil {
		return authcoremodel.LoginResponse{}, err
	}
	log.Printf("User with email %s logged in", request.Email)

	token := createToken(authUser.UserId, authUser.Email)
	return authcoremodel.LoginResponse{
		UserId:   authUser.UserId,
		Username: authUser.Username,
		Email:    authUser.Email,
		Token:    token,
	}, nil
}

//func ResetPassword(user authcoremodel.User) error {
//
//}

func createToken(userId string, email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"iss":    issuer,
		"aud":    audience,
		"exp":    time.Now().Add(time.Minute * 20).Unix(),
	})
	tokenString, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Generated token for user with email %s", email)
	return tokenString
}

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Println("Authorization header blank")
			c.Status(403)
			c.Abort()
		}

		tokenParts := strings.Split(authHeader, " ")

		token, err := jwt.Parse(tokenParts[1], validateToken)

		log.Printf("Token, error")
		log.Println(token)
		log.Println(err)

		if token.Valid {
			log.Println("Token valid")
			c.Next()
		}
		c.Status(403)
		c.Abort()
	}
}

func validateToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("invalid signing method")
	}

	if !token.Claims.(jwt.MapClaims).VerifyAudience(audience, true) {
		return nil, fmt.Errorf("invalid audience")
	}

	if !token.Claims.(jwt.MapClaims).VerifyIssuer(issuer, true) {
		return nil, fmt.Errorf("invalid issuer")
	}
	return mySigningKey, nil
}
