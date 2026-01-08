package database

import (
	"log"
	"os"
	"time"

	"mamajulia/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false, // Desabilita prepared statements para evitar conflitos
	})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	// Configura o pool de conexões do SQL subjacente
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Erro ao obter conexão SQL:", err)
	}

	// Configurações do pool de conexões
	sqlDB.SetMaxIdleConns(10)                  // Máximo de conexões idle
	sqlDB.SetMaxOpenConns(100)                 // Máximo de conexões abertas
	sqlDB.SetConnMaxLifetime(time.Hour)        // Tempo máximo de vida da conexão
	sqlDB.SetConnMaxIdleTime(10 * time.Minute) // Tempo máximo que uma conexão pode ficar idle

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
