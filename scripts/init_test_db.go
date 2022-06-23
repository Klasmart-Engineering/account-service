package main

import (
	"fmt"
	"io/ioutil"
	db "kidsloop/account-service/database"
	"kidsloop/account-service/test_util"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Initializing test database")

	baseDir := "./database/migrations/"
	sqlMigrations := make([]string, 0)

	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			content, err := ioutil.ReadFile(fmt.Sprintf("%s%s", baseDir, file.Name()))
			if err != nil {
				log.Fatal(err)
			}

			sqlMigrations = append(sqlMigrations, string(content))
		}
	}

	test_util.LoadTestEnv("./")
	err = db.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to postgres")
	}

	for _, sqlMigration := range sqlMigrations {
		fmt.Println(sqlMigration)
		_, err := db.Database.Conn.Exec(sqlMigration)
		if err != nil {
			log.Fatal(err)
		}

	}

	fmt.Println("Database initialization complete")
}
