package main

import (
	"log"

	"temp/internal/encrypt"
	"temp/internal/models"
	"temp/internal/users"
	"temp/internal/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Replace with your PostgreSQL connection string.
	db, err := utils.LoadDb()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	userHandler := &users.UserHandler{DB: db}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // Allows all origins (not recommended for production)
	router.Use(cors.New(config))
	router.Use(utils.SanitizeMiddleware)
	userHandler.SetupRoutes(router)
	encrypt.SetupRoutes(router)
	port := ":8080"
	log.Printf("Server listening on port %s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
