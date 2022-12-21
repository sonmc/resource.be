package authController

import (
	"net/http"

	authDto "resource_be/application/dto"

	"github.com/gin-gonic/gin"
)


func HandleLogin( ) gin.HandlerFunc {

	//var dbConn *gorm.DB = mysqlDb.SetupDatabaseConnection()
	return func(c *gin.Context) {
		var auth authDto.Auth

		if err := c.ShouldBind(&auth); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
  
		// if err := auth.Validate(); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		//storage := authRepository.NewMySQLStorage(dbConn);
		//biz := authBiz.LoginBiz(storage);

		// if err := biz.OnLogin(c.Request.Context(), &auth); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		c.JSON(http.StatusOK, gin.H{"data": "auth.Id"})
	}
}
