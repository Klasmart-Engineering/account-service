package db

import (
	"context"
	"database/sql"
	"fmt"

	util "kidsloop/account-service/util"

	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
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

	connection, err := sql.Open("nrpostgres", connStr)
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

func RunInTransaction[T any](ctx context.Context, handler func(tx *sql.Tx) (T, error)) (T, error) {
	var result T

	tx, err := Database.Conn.BeginTx(ctx, nil)
	if err != nil {
		return result, err
	}
	defer tx.Rollback()

	result, err = handler(tx)
	if err != nil {
		return result, err
	}

	if err = tx.Commit(); err != nil {
		return result, err
	}

	return result, nil
}
