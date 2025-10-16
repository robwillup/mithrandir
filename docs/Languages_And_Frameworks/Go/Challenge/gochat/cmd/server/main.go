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
	"gochat/internal/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "postgres://admin:secret@localhost:5432/gochat?sslmode=disable" //os.Getenv("DATABASE_URL")
	print(dsn)
	db.Init(dsn)

	r := gin.Default()
	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)

	log.Fatal(r.Run(":8080"))
}
