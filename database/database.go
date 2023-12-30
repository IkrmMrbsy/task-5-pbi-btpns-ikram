
package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"PBIP_BTPN_SYARIAH/models"
)

var (
	DB *gorm.DB
)

func InitDB() {
	var err error
	dsn := "appuser@tcp(127.0.0.1:3306)/photoapp_api?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Failed to connect to database. Error: %v\n", err)
		panic("Failed to connect to database")
	}

	DB.AutoMigrate(&models.User{}, &models.Photo{})
}
