package routes

import (
	controller "be/application/controllers"
	service "be/application/services"
	"net/http"
)

var (  
	authService       = service.NewAuthService(userRepository)
 	authController    = controller.NewAuthController(logger, authService) 
)

func authRoutes() {
	httpRouter.GET("/auth", func(response http.ResponseWriter, request *http.Request) {
		logger.Println(response, "Up and running...")
	}) 
	httpRouter.POST("/login", authController.Login)
}