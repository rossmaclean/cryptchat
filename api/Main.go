package main

import (
	authleft "cryptchat/auth/left"
	chatsleft "cryptchat/chats/left"
	properties2 "cryptchat/properties"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	log.Println("Starting Application")
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	if properties2.GetEnv() != "local" {
		router.LoadHTMLGlob("../frontend/build/*.html")
		router.Static("/static", "./frontend/build/static")
		router.GET("/", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{})
		})
	}

	//router.Use(IsAuth())

	router.POST("/api/v1/signup", authleft.SignupHandler)
	router.POST("/api/v1/login", authleft.LoginHandler)

	router.POST("/api/v1/chats", IsAuth(), chatsleft.ChatsHandler)
	router.POST("/api/v1/chat_messages", IsAuth(), chatsleft.ChatMessageHandler)

	err := router.Run(":8000")
	if err != nil {
		log.Fatal("Unable to start web server", err)
	}
	log.Println("Application Running")
}

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			tk := c.GetHeader("Authorization")
			log.Println("Token: " + tk)
			sp := strings.Split(tk, " ")
			log.Println(sp)

			token, err := jwt.Parse(sp[1], func(token *jwt.Token) (interface{}, error) {
				token.Valid = false
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Invalid Signing Method")
				}
				aud := "billing.jwtgo.io"
				checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
				if !checkAudience {
					return nil, fmt.Errorf("invalid aud")
				}
				// verify iss claim
				// BUG TODO - When field was set as jss accidentally in token,
				// this returned true
				iss := "cryptchat.rossmac.co.uk"
				checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
				log.Println(checkIss)
				if !checkIss {
					return nil, fmt.Errorf("invalid iss")
				}

				token.Valid = true
				return MySigningKey, nil
			})
			log.Println("Token:")
			log.Println(token)
			if err != nil {
				log.Println(err.Error())
			}

			if token.Valid {
				log.Println("Token Valid")
				c.Next()
			}

		} else {
			//fmt.Fprintf(w, "No Authorization Token provided")
			log.Println("No Auth token provided")
			c.Abort()
		}
	}
}
