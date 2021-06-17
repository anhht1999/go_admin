package database

import (
	"github.com/anhht1999/go_admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Connect()  {
	dsn := "root:@tcp(127.0.0.1:3306)/go_admin?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}
	// fmt.Println(db)

	DB = database

	database.AutoMigrate(&models.User{})
}