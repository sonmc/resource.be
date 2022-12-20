package migrations

import (
	mysqlDb "be/database"
	"be/entities"

	"gorm.io/gorm"
)

func Migrate() {
	var dbConn *gorm.DB = mysqlDb.SetupDatabaseConnection()
	dbConn.AutoMigrate(entities.User{})

	// dbConn.Model(&models.User{}).AddForeignKey("address_id", "address(id)", "CASCADE", "RESTRICT")
	// dbConn.Model(&models.Orders{}).AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT")
	 
}
