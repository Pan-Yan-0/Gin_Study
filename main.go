// main.go
package main

import (
	"LoginStudy/app/database"
	"LoginStudy/app/route"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	database.InitializeDB()

	// Create a Gin router
	router := gin.Default()

	// Set up routes
	route.SetupRoutes(router)

	// Start the Gin server
	router.Run(":8080")
}
