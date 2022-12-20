package routes

import (
	authTransport "be/modules/auth/transport"

	"github.com/gin-gonic/gin"
)

func authRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth") 
	auth.POST("/login", authTransport.HandleLogin()) 
}