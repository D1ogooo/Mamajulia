package database

import (
	"log"
	"mamajulia/src/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("mamajulia.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	db.AutoMigrate(&models.User{}, &models.Dish{})
	DB = db
}
