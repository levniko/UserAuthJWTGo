package database

import (
	"test3/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var psqlconn string = "host=localhost user=postgres password=1573596248 dbname=go_auth sslmode=disable"

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	
	DB = connection

	connection.AutoMigrate(&models.User{})
}
