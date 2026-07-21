package main

import (
	"ginlearn/internal/config"
	"ginlearn/internal/db"
	"ginlearn/internal/server"
	"log"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Config Error: %v", err)
	}

	pool, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}
	defer db.Disconnect(pool)

	log.Println("Connected to PostgreSQL")

	router := server.NewRouter(pool)

	log.Printf("Server running on :%s", cfg.ServerPort)

	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
