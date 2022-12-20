package authController

import (
	database "be/database"
	authBiz "be/modules/auth/business"
	authDto "be/modules/auth/dto"
	authStorage "be/modules/auth/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func HandleLogin( ) gin.HandlerFunc {

	var dbConn *gorm.DB = database.SetupDatabaseConnection()
	return func(c *gin.Context) {
		var auth authDto.Auth

		if err := c.ShouldBind(&auth); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
  
		if err := auth.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := authStorage.NewMySQLStorage(dbConn);
		biz := authBiz.LoginBiz(storage);

		if err := biz.OnLogin(c.Request.Context(), &auth); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": "auth.Id"})
	}
}
