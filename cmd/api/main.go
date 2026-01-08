package main

import (
	"log"
	"mamajulia/internal/database"
	"mamajulia/internal/routes"
	config "mamajulia/pkg/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.ConnectDatabase()

	r := gin.Default()
	routes.SetupRoutes(r)

	log.Printf("Servidor rodando na porta 3000 🚀")
	r.Run(":3000")
}