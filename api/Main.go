package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting Application")
	router := gin.Default()
	router.LoadHTMLGlob("frontend/build/*.html")
	router.Static("/static", "./frontend/build/static")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.Run(":8000")
	log.Println("Application Running")
}
