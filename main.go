package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"temp/internal/auth"
	"temp/internal/encrypt"
	"temp/internal/models"
	"temp/internal/users"
	"temp/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Replace with your PostgreSQL connection string.
	databaseHost := os.Getenv("POSTGRES_HOST")
	databasePort := os.Getenv("POSTGRES_PORT")
	databaseDb := os.Getenv("POSTGRES_DB")
	databasePassword := os.Getenv("POSTGRES_PASSWORD")
	databaseUser := os.Getenv("POSTGRES_USERNAME")
	fmt.Println("database info", databaseHost, databasePort, databaseUser, databaseDb, databasePassword)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Berlin", databaseHost, databaseUser, databasePassword, databaseDb, databasePort)

	// Connect to the database (PostgreSQL in this example).
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
	router.Use(utils.SanitizeMiddleware)
	router.POST("/login", userHandler.LoginUser)
	router.POST("/register", userHandler.RegisterUser)
	protected := router.Group("/protected")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("/resource", func(c *gin.Context) {
			claims, _ := c.Get("claims")
			c.JSON(http.StatusOK, gin.H{"message": "Protected resource accessed", "claims": claims})
		})
		protected.POST("/encrypt", encrypt.EncryptHandler)
		protected.POST("/decrypt", encrypt.DecryptHandler)
		protected.POST("/hash", encrypt.HashingHandler)
	}

	port := ":8080"
	log.Printf("Server listening on port %s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
