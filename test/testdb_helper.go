package test

import (
	"fmt"
	"log"
	"os"

	"github.com/alvinmdj/mygram-api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func loadTestEnv() {
	// load env with relative path to the .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file => ", err)
	}
}

func SetupTestDB() *gorm.DB {
	loadTestEnv()

	dbHost := os.Getenv("TEST_DB_HOST")
	dbPort := os.Getenv("TEST_DB_PORT")
	dbUser := os.Getenv("TEST_DB_USER")
	dbPassword := os.Getenv("TEST_DB_PASSWORD")
	dbName := os.Getenv("TEST_DB_NAME")

	config := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)
	dsn := config
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database:", err.Error())
	}

	log.Println("database connected successfully")

	// auto migrate table schemas
	db.AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})

	return db
}

// Truncates the users table in the given database.
func TruncateUsersTable(db *gorm.DB) {
	if err := db.Exec("TRUNCATE TABLE users CASCADE").Error; err != nil {
		log.Fatalf("Error truncating users table: %v", err)
	}
}
