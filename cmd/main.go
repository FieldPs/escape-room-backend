package main

import (
	"fmt"
	"os"

	"github.com/FieldPs/escape-room-backend/internal/models"
	"github.com/FieldPs/escape-room-backend/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: Could not load .env file")
	}

	// Construct PostgreSQL DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Disables all SQL logging
	})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	db.AutoMigrate(&models.User{}, &models.Puzzle{}, &models.UserPuzzle{})

	// Set up Gin router
	r := gin.Default()
	routes.SetupRoutes(r, db)

	// Run server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
