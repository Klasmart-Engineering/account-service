package db

import (
	"database/sql"
	"kidsloop/account-service/model"
)

func (db DB) CreateAccount(tx *sql.Tx) (model.Account, error) {
	sql := `INSERT INTO account DEFAULT VALUES RETURNING id`
	account := model.Account{}

	var err error
	if tx != nil {
		err = tx.QueryRow(sql).Scan(&account.ID)
	} else {
		err = db.Conn.QueryRow(sql).Scan(&account.ID)
	}

	return account, err
}

func (db DB) GetAccount(tx *sql.Tx, id string) (model.Account, error) {
	sql := `SELECT id FROM account WHERE id = $1 LIMIT 1`
	account := model.Account{}

	var err error
	if tx != nil {
		err = tx.QueryRow(sql, id).Scan(&account.ID)
	} else {
		err = db.Conn.QueryRow(sql, id).Scan(&account.ID)
	}

	return account, err
}

func (db DB) DeleteAccount(tx *sql.Tx, id string) error {
	sql := `DELETE FROM account WHERE id = $1`

	var err error
	if tx != nil {
		_, err = tx.Exec(sql, id)
	} else {
		_, err = db.Conn.Exec(sql, id)
	}

	return err
}
