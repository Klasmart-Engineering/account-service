package main

import (
	db "kidsloop/account-service/database"
	_ "kidsloop/account-service/docs"
	"kidsloop/account-service/handler"
	_ "kidsloop/account-service/util"
	"log"
	"os"

	"github.com/joho/godotenv"
	nrgin "github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
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

	// Create New Relic agent ("Application")
	nrApp, nrErr := newrelic.NewApplication(
		newrelic.ConfigAppName("New Relic Monitoring"),
		newrelic.ConfigLicense("__NEW_RELIC_LICENSE_KEY"),
		newrelic.ConfigDebugLogger(os.Stdout),
	)
	if nrErr != nil {
		log.Println("Unable to create New Relic Monitoring. Reason:", nrErr)
	}

	router := handler.SetUpRouter()
	router.Use(nrgin.Middleware(nrApp))
	router.Run()

	log.Println("Started router")
}
