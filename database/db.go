package db

import (
	"database/sql"
	"fmt"

	util "kidsloop/account-service/util"

	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

var Database DB

func InitDB() error {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		util.GetEnvOrPanic("POSTGRES_HOST"),
		util.GetEnvOrPanic("POSTGRES_PORT"),
		util.GetEnvOrPanic("POSTGRES_USER"),
		util.GetEnvOrPanic("POSTGRES_PASSWORD"),
		util.GetEnvOrPanic("POSTGRES_DB"),
	)

	connection, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = connection.Ping()
	if err != nil {
		return err
	}

	Database = DB{}
	Database.Conn = connection

	return nil
}
