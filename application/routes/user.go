package routes

import (
	authController "resource_be/application/controllers"

	"github.com/gin-gonic/gin"
)

func userRoutes(rg *gin.RouterGroup) { 

	users := rg.Group("/users")
	users.GET("/", authController.HandleLogin( ))
	users.POST("", authController.HandleLogin( ))
 
}
