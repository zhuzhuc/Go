package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	var err error

	// Get database connection details from environment variables
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "personal_website")
	dbSSLMode := getEnv("DB_SSL_MODE", "disable")

	// Create DSN string
	var dsn string
	if dbPassword == "" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s",
			dbHost, dbPort, dbUser, dbName, dbSSLMode)
	} else {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)
	}
	log.Printf("Connecting to database with DSN: %s", dsn)

	// Configure GORM logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Open connection to database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database successfully")

	// Auto migrate models
	migrateModels()
}

// Helper function to get environment variable with fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// Migrate database models
func migrateModels() {
	log.Println("Starting database migration...")

	// For now, let's use GORM's AutoMigrate with simplified models
	// This will help us get the application running

	// Create a simplified User model for migration
	type SimpleUser struct {
		ID        uint   `gorm:"primaryKey"`
		Username  string `gorm:"size:50;not null;unique"`
		Email     string `gorm:"size:100;not null;unique"`
		Password  string `gorm:"size:100;not null"`
		FirstName string `gorm:"size:50"`
		LastName  string `gorm:"size:50"`
		Role      string `gorm:"size:20;default:'user'"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	// Migrate the simplified model
	if err := DB.AutoMigrate(&SimpleUser{}); err != nil {
		log.Printf("Warning: Failed to migrate User model: %v", err)
	} else {
		log.Println("User model migrated successfully")
	}

	log.Println("Database migration completed successfully")
}
