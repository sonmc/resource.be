package authService

import (
	"context"
	authDto "resource_be/application/dto"
)

type LoginStorage interface {
	Login(ctx context.Context, data *authDto.Auth)  error
}

type loginBiz struct {
	store LoginStorage
}

func LoginBiz(store LoginStorage) *loginBiz {
	return &loginBiz{store: store}
}

func (biz *loginBiz) OnLogin(ctx context.Context, data *authDto.Auth) error {
	if err := biz.store.Login(ctx, data); err != nil {
		return err
	} 
	return nil
}
