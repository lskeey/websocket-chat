package main

import (
	"log"

	"github.com/lskeey/websocket-chat/config"
	"github.com/lskeey/websocket-chat/models"
)

func main() {
	db := config.ConnectDB()

	err := db.AutoMigrate(
		&models.User{},
		&models.Message{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrated successfully!")
}
