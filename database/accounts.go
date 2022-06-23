package db

import (
	"database/sql"
	"errors"
	"fmt"
	api_errors "kidsloop/account-service/errors"
	"kidsloop/account-service/model"
	"net/http"
)

func (db DB) CreateAccount(tx *sql.Tx) (model.Account, error) {
	query := `INSERT INTO account DEFAULT VALUES RETURNING id`
	account := model.Account{}

	var err error
	if tx != nil {
		err = tx.QueryRow(query).Scan(&account.ID)
	} else {
		err = db.Conn.QueryRow(query).Scan(&account.ID)
	}

	return account, err
}

func (db DB) GetAccount(tx *sql.Tx, id string) (model.Account, error) {
	query := `SELECT id FROM account WHERE id = $1 LIMIT 1`
	account := model.Account{}

	var err error
	if tx != nil {
		err = tx.QueryRow(query, id).Scan(&account.ID)
	} else {
		err = db.Conn.QueryRow(query, id).Scan(&account.ID)
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

func (db DB) DeleteAccount(tx *sql.Tx, id string) error {
	query := `DELETE FROM account WHERE id = $1`

	var result sql.Result
	var err error
	if tx != nil {
		result, err = tx.Exec(query, id)
	} else {
		result, err = db.Conn.Exec(query, id)
	}

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err == nil && rowsAffected == 0 {
		return &api_errors.APIError{
			Status:  http.StatusNotFound,
			Code:    api_errors.ErrCodeNotFound,
			Message: fmt.Sprintf(api_errors.ErrMsgNotFound, "account", id),
			Err:     errors.New("no rows affected"),
		}
	}

	return err
}
