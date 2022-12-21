package routes

import (
	authController "resource_be/application/controllers"

	"github.com/gin-gonic/gin"
)

func authRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth") 
	auth.POST("/login", authController.HandleLogin()) 
}