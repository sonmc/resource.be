package service

import (
	repository "be/application/repositories"

	model "be/domain/user"
)

type AuthService interface {
	Login() (model.User, error)
}

type authService struct{
	repository repository.UserRepository
}


func NewAuthService(repository repository.UserRepository) AuthService {
	return &authService{repository: repository}
}

func (service *authService) Login() (model.User, error) {
	return  model.User{}, nil;
}
