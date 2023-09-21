package test

import (
	"fmt"
	"log"
	"os"

	"github.com/alvinmdj/mygram-api/models"
	"github.com/alvinmdj/mygram-api/repositories"
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

// Truncates the social media table in the given database.
func TruncateSocialMediasTable(db *gorm.DB) {
	if err := db.Exec("TRUNCATE TABLE social_media CASCADE").Error; err != nil {
		log.Fatalf("Error truncating social_media table: %v", err)
	}
}

// function to create a user before login test
func RegisterTestUser(db *gorm.DB) models.User {
	userData := models.User{
		Username: "alvinmdj",
		Email:    "alvinmdj@mygram.com",
		Password: "password",
		Age:      20,
	}
	userRepo := repositories.NewUserRepo(db)
	result, _ := userRepo.Save(userData)
	return result
}

func GetTestUser(db *gorm.DB) (models.User, error) {
	userData := models.User{
		Email: "alvinmdj@mygram.com",
	}
	userRepo := repositories.NewUserRepo(db)
	return userRepo.FindByEmail(userData)
}

func CreateTestSocialMedia(db *gorm.DB) models.SocialMedia {
	userData, _ := GetTestUser(db)

	socialMediaData := models.SocialMedia{
		UserID:         userData.ID,
		Name:           "My Youtube",
		SocialMediaURL: "https://www.youtube.com",
	}

	socialMediaRepo := repositories.NewSocialMediaRepo(db)
	result, _ := socialMediaRepo.Save(socialMediaData)
	return result
}
