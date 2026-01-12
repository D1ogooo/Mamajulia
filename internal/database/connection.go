package database

import (
	"log"
	"os"
	"strings"
	"time"

	"mamajulia/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")

	if dsn != "" && !strings.Contains(dsn, "prefer_simple_protocol") {
		if strings.Contains(dsn, "?") {
			dsn += "&prefer_simple_protocol=true"
		} else {
			dsn += "?prefer_simple_protocol=true"
		}
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Erro ao obter conexão SQL:", err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	err = db.AutoMigrate(
		&models.User{},
		&models.Dish{},
		&models.Order{},
		&models.OrderDish{},
	)

	if err != nil {
		log.Fatal("Erro ao migrar tabelas:", err)
	}

	DB = db
	log.Println("✅ Conectado ao Neon com sucesso!")
}
