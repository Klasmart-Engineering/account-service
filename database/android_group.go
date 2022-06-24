package db

import (
	"database/sql"
	"fmt"
	api_errors "kidsloop/account-service/errors"
	"kidsloop/account-service/model"
	"net/http"
)

func (db DB) CreateAndroidGroup(tx *sql.Tx, accountID string) (model.AndroidGroup, error) {
	query := `INSERT INTO android_group (account_id) VALUES ($1) RETURNING id`
	androidGroup := model.AndroidGroup{}

	var err error
	if tx != nil {
		err = tx.QueryRow(query, accountID).Scan(&androidGroup.ID)
	} else {
		err = db.Conn.QueryRow(query, accountID).Scan(&androidGroup.ID)
	}

	return androidGroup, err
}

func (db DB) GetAndroidGroup(tx *sql.Tx, id string) (model.AndroidGroup, error) {
	query := `SELECT id, account_id FROM android_group WHERE id = $1 LIMIT 1`
	androidGroup := model.AndroidGroup{}

	var err error
	if tx != nil {
		err = tx.QueryRow(query, id).Scan(&androidGroup.ID, &androidGroup.ID)
	} else {
		err = db.Conn.QueryRow(query, id).Scan(&androidGroup.ID, &androidGroup.AccountID)
	}

	if err == sql.ErrNoRows {
		return androidGroup, &api_errors.APIError{
			Status:  http.StatusNotFound,
			Code:    api_errors.ErrCodeNotFound,
			Message: fmt.Sprintf(api_errors.ErrMsgNotFound, "android_group", id),
			Err:     err,
		}
	}

	return androidGroup, err
}
