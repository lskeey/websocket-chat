package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lskeey/websocket-chat/config"
	"github.com/lskeey/websocket-chat/handlers"
	"github.com/lskeey/websocket-chat/models"
)

func main() {
	DB := config.ConnectDB()

	err := DB.AutoMigrate(
		&models.User{},
		&models.Message{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrated successfully!")

	router := gin.Default()

	authHandler := &handlers.AuthHandler{DB: DB}

	api := router.Group("/api")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
