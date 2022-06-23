package db

import (
	"database/sql"
	"kidsloop/account-service/model"
)

func (db DB) CreateAndroidGroup(tx *sql.Tx, accountID string) (model.AndroidGroup, error) {
	sql := `INSERT INTO android_group (account_id) VALUES ($1) RETURNING id`
	androidGroup := model.AndroidGroup{}

	var err error
	if tx != nil {
		err = tx.QueryRow(sql, accountID).Scan(&androidGroup.ID)
	} else {
		err = db.Conn.QueryRow(sql, accountID).Scan(&androidGroup.ID)
	}

	return androidGroup, err
}
