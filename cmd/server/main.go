package main

import (
	"log"

	"payment-gateway/internal/api"
	"payment-gateway/internal/config"
	"payment-gateway/internal/database"
)

func main() {
	cfg := config.LoadConfig()

	db := database.NewPostgres(cfg.DBUrl)

	router := api.SetupRouter(cfg, db)

	log.Println("Server running on port", cfg.Port)

	router.Run(":" + cfg.Port)

}
