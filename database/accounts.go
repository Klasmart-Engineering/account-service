package db

import (
	"context"
	"database/sql"
	"fmt"
	api_errors "kidsloop/account-service/errors"
	"kidsloop/account-service/model"
	"net/http"
)

func (db DB) CreateAccount(tx *sql.Tx, ctx context.Context) (model.Account, error) {
	query := `INSERT INTO account DEFAULT VALUES RETURNING id`
	account := model.Account{}

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query).Scan(&account.ID)
	} else {
		err = db.Conn.QueryRowContext(ctx, query).Scan(&account.ID)
	}

	return account, err
}

func (db DB) GetAccount(tx *sql.Tx, id string, ctx context.Context) (model.Account, error) {
	query := `SELECT id FROM account WHERE id = $1 LIMIT 1`
	account := model.Account{}

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, id).Scan(&account.ID)
	} else {
		err = db.Conn.QueryRowContext(ctx, query, id).Scan(&account.ID)
	}

	if err == sql.ErrNoRows {
		return account, &api_errors.APIError{
			Status:  http.StatusNotFound,
			Code:    api_errors.ErrCodeNotFound,
			Message: fmt.Sprintf(api_errors.ErrMsgNotFound, "account", id),
			Err:     err,
		}
	}

	return account, err
}

func (db DB) DeleteAccount(tx *sql.Tx, id string, ctx context.Context) error {
	query := `DELETE FROM account WHERE id = $1 RETURNING id`
	var accountId string

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, id).Scan(&accountId)
	} else {
		err = db.Conn.QueryRowContext(ctx, query, id).Scan(&accountId)
	}

	if err == sql.ErrNoRows {
		return &api_errors.APIError{
			Status:  http.StatusNotFound,
			Code:    api_errors.ErrCodeNotFound,
			Message: fmt.Sprintf(api_errors.ErrMsgNotFound, "account", id),
			Err:     err,
		}
	}

	return err
}
