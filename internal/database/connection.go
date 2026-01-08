package database

import (
	"log"
	"mamajulia/internal/models"
	config "mamajulia/pkg/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=" + config.Get("DB_HOST") +
		" user=" + config.Get("DB_USER") +
		" password=" + config.Get("DB_PASSWORD") +
		" dbname=" + config.Get("DB_NAME") +
		" port=" + config.Get("DB_PORT") +
		" sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	db.AutoMigrate(&models.User{}, &models.Dish{}, &models.Order{})
	DB = db

	log.Println("✅ Conectado ao PostgreSQL com sucesso!")
}