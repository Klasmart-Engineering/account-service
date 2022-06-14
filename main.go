package main

import (
	db "kidsloop/account-service/database"
	_ "kidsloop/account-service/docs"
	"kidsloop/account-service/handler"
	_ "kidsloop/account-service/util"
	"log"

	"github.com/joho/godotenv"
)

// @title    account-service documentation
// @version  0.0.1
// @host     localhost:8080

func main() {
	log.Println("Starting account-service")

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: no .env file found.")
	}

	err = db.InitDB()
	if err != nil {
		log.Println("Failed to connect to postgres:")
		log.Fatal(err)
	}

	log.Println("Connected to Postgres")

	router := handler.SetUpRouter()
	router.Run()

	log.Println("Started router")
}
