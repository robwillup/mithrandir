/*
* Initialize DB connection
* Setup Gin router
* Register authentication routes from internal/auth/handler.go
* Add middleware for JWT
* Run server
 */

package main

import (
	"gochat/internal/auth"
	"gochat/internal/chat"
	"gochat/internal/db"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	db.Init(dsn)

	r := gin.Default()
	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)

	authorized := r.Group("/")
	authorized.Use(auth.JWTAuthMiddleware())

	authorized.POST("/messages", chat.PostMessage)
	authorized.GET("/messages", chat.GetMessages)

	go HandleMessages()
	r.GET("/ws", HandleConnections)

	log.Fatal(r.Run(":8080"))
}
