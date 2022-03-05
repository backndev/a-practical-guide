package db

import (
	"go-admin/main/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:@/go_admin"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}

	DB = db

	_ = db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
}
