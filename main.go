package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lskeey/websocket-chat/config"
	"github.com/lskeey/websocket-chat/handlers"
	"github.com/lskeey/websocket-chat/middleware"
	"github.com/lskeey/websocket-chat/models"
	"github.com/lskeey/websocket-chat/ws"
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

	hub := ws.NewHub(DB)
	go hub.Run()

	router := gin.Default()

	config := cors.DefaultConfig()
  config.AllowOrigins = []string{"http://localhost:3000"}
  config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
  config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
  config.ExposeHeaders = []string{"Content-Length"}
  config.AllowCredentials = true
  config.MaxAge = 12 * 3600

  router.Use(cors.New(config))

	authHandler := &handlers.AuthHandler{DB: DB}
	userHandler := &handlers.UserHandler{DB: DB}
	messageHandler := &handlers.MessageHandler{DB: DB}
	websocketHandler := &handlers.WebsocketHandler{Hub: hub}

	api := router.Group("/api")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)
	}

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users/search", userHandler.SearchUsers)
		protected.GET("/messages/:recipient_id", messageHandler.GetMessages)

		protected.GET("/ws", websocketHandler.ServeWs)
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
