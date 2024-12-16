// app/controller/LoginController.go
package controller

import (
	"LoginStudy/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginRequest struct to bind the JSON request body
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login handles user login
func Login(c *gin.Context) {
	var loginReq LoginRequest

	// Bind JSON to struct
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Call the LoginService to authenticate
	if err := service.Authenticate(loginReq.Username, loginReq.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
