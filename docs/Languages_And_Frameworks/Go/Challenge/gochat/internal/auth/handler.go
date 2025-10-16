/*
* Gin HTTP handlers for:
*  - POST /register to create user with hashed password.
*  - POST /login to verify password, generate, and return JWT token.
*/

package auth

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gochat/internal/db"
	"gochat/internal/models"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"username" binding:"required"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	passwordHash, err := HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
	}

	_, err = db.DB.Exec(`INSERT INTO users (username, password_hash) VALUES ($1, $2)`, req.Username, passwordHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username might be taken"})
		return
	}

	c.JSON(http)
}
