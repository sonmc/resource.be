package authStorage

import (
	"be/entities"
	authdto "be/modules/auth/dto"
	"context"

	"gorm.io/gorm"
)
 
func (s *mysqlStorage) Login( ctx context.Context,  auth *authdto.Auth)  ( error) {
 
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
