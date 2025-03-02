package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDb() (*gorm.DB, error) {
	databaseHost := os.Getenv("POSTGRES_HOST")
	databasePort := os.Getenv("POSTGRES_PORT")
	databaseDb := os.Getenv("POSTGRES_DB")
	databasePassword := os.Getenv("POSTGRES_PASSWORD")
	databaseUser := os.Getenv("POSTGRES_USERNAME")
	fmt.Println("database info", databaseHost, databasePort, databaseUser, databaseDb, databasePassword)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Berlin", databaseHost, databaseUser, databasePassword, databaseDb, databasePort)

	// Connect to the database (PostgreSQL in this example).
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
