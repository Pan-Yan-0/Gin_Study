// app/controller/LoginController.go
package controller

import (
	"LoginStudy/app/repository"
	"LoginStudy/app/service"
	"LoginStudy/app/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginRequest struct to bind the JSON request body
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login handles user login and returns JWT if credentials are valid
func Login(c *gin.Context) {
	var loginReq LoginRequest

	// Bind JSON to struct
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Instantiate the service and authenticate the user
	loginService := service.LoginService{UserRepo: &repository.UserRepository{}}
	if err := loginService.Authenticate(loginReq.Username, loginReq.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT token
	token, err := util.GenerateToken(loginReq.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Respond with the token
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
