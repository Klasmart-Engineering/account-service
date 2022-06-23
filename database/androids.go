package db

import (
	"database/sql"
	"kidsloop/account-service/model"
)

func (db DB) CreateAndroid(tx *sql.Tx, accountID string, androidGroupId string) (model.Android, error) {
	sql := `INSERT INTO android (android_group_id) VALUES ($1) RETURNING id, android_group_id`
	android := model.Android{}

	var err error
	if tx != nil {
		err = tx.QueryRow(sql, androidGroupId).Scan(&android.ID, &android.AndroidGroupID)
	} else {
		err = db.Conn.QueryRow(sql, androidGroupId).Scan(&android.ID, &android.AndroidGroupID)
	}

	return android, err
}

func (db DB) GetAndroid(tx *sql.Tx) (model.Android, error) {
	sql := `SELECT id, android_group_id FROM android WHERE id = $1 LIMIT 1`
	android := model.Android{}

	var err error
	if tx != nil {
		err = tx.QueryRow(sql, android.ID).Scan(&android.ID, &android.AndroidGroupID)
	} else {
		err = db.Conn.QueryRow(sql, android.ID).Scan(&android.ID, &android.AndroidGroupID)
	}

	return android, err
}

func (db DB) DeleteAndroid(tx *sql.Tx, id string) error {
	sql := `DELETE FROM android WHERE id = $1`
	var err error
	if tx != nil {
		_, err = tx.Exec(sql, id)
	} else {
		_, err = db.Conn.Exec(sql, id)
	}

	return err
}
