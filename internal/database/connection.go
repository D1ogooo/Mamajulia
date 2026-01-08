package database

import (
	"log"
	"os"

	"mamajulia/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Dish{},
		&models.Order{},
	)
	if err != nil {
		log.Fatal("Erro ao migrar tabelas:", err)
	}

	DB = db
	log.Println("✅ Conectado ao Neon com sucesso!")
}
