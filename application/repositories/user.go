package authStorage

import (
	"context"

	authDto "resource_be/application/dto"

	"resource_be/domain/entities"

	"gorm.io/gorm"
)
 
func (s *mysqlStorage) Login( ctx context.Context,  auth *authDto.Auth)  ( error) {
 
	if err := s.db.Table(entities.User{}.TableName()).
		Where("username=? and password=?", auth.Username, auth.Password).
		First(&auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound {  
			return nil 
		}
		return nil 
	} 
	return nil
}
