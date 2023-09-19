package database

import (
	"fmt"
	"log"
	"os"

	"github.com/alvinmdj/mygram-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	config := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database:", err.Error())
	}

	log.Println("database connected successfully")

	// auto migrate table schemas
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
}

// variadic param to get testDB for go test
func GetDB(testDB ...*gorm.DB) *gorm.DB {
	if len(testDB) > 0 && testDB[0] != nil {
		return testDB[0]
	}
	return db
}
