package main

import (
	"log"

	"github.com/alvinmdj/mygram-api/database"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	database.StartDB()
}
