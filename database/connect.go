package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect()  {
	dsn := "root:@tcp(127.0.0.1:3306)/go_admin?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}
	// fmt.Println(db)
}