package routes

import (
	authTransport "be/modules/auth/transport"

	"github.com/gin-gonic/gin"
)

func userRoutes(rg *gin.RouterGroup) { 

	users := rg.Group("/users")
	users.GET("/", authTransport.HandleLogin( ))
	users.POST("", authTransport.HandleLogin( ))
 
}
