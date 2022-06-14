package db

import (
	"kidsloop/account-service/model"
)

func (db DB) CreateAccount() (model.Account, error) {
	sql := `INSERT INTO account DEFAULT VALUES RETURNING id`
	account := model.Account{}
	err := db.Conn.QueryRow(sql).Scan(&account.ID)
	return account, err
}

func (db DB) GetAccount(id string) (model.Account, error) {
	sql := `SELECT id FROM account WHERE id = $1 LIMIT 1`
	account := model.Account{}
	err := db.Conn.QueryRow(sql, id).Scan(&account.ID)
	return account, err
}

func (db DB) DeleteAccount(id string) error {
	sql := `DELETE FROM account WHERE id = $1`
	_, err := db.Conn.Exec(sql, id)
	return err
}
