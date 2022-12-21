package controller

import (
	"encoding/json"
	"log"
	"net/http"

	service "be/application/services"
) 

type AuthController interface {
	Login(response http.ResponseWriter, request *http.Request)
}

type authController struct{
	authService service.AuthService
	logger *log.Logger
}

func NewAuthController(logger *log.Logger, authService service.AuthService) AuthController {
	return &authController{authService: authService, logger: logger}
}

func (auth *authController) Login(response http.ResponseWriter, request *http.Request) {

	auth.logger.Println("GetProducts controller method called ")

	response.Header().Set("Content-Type", "application/json")
	products, err := auth.authService.Login()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Internal server error"}`))
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(products)
}
