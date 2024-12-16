package route

import (
	"LoginStudy/app/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	userRoutes := router.Group("/user")
	{
		// POST route for login
		userRoutes.POST("/login", controller.Login) // POST /user/login
	}

}
