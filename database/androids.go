package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	api_errors "kidsloop/account-service/errors"
	"kidsloop/account-service/model"
	"net/http"
)

func (db DB) CreateAndroid(tx *sql.Tx, ctx context.Context, accountID string, androidGroupId string) (model.Android, error) {
	query := `INSERT INTO android (android_group_id) VALUES ($1) RETURNING id, android_group_id`
	android := model.Android{}

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, androidGroupId).Scan(&android.ID, &android.AndroidGroupID)
	} else {
		err = db.Conn.QueryRowContext(ctx, query, androidGroupId).Scan(&android.ID, &android.AndroidGroupID)
	}

	return android, err
}

func (db DB) GetAndroid(tx *sql.Tx, ctx context.Context, id string) (model.Android, error) {
	query := `SELECT id, android_group_id FROM android WHERE id = $1 LIMIT 1`
	android := model.Android{}

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, id).Scan(&android.ID, &android.AndroidGroupID)
	} else {
		err = db.Conn.QueryRowContext(ctx, query, id).Scan(&android.ID, &android.AndroidGroupID)
	}

	if err == sql.ErrNoRows {
		return android, &api_errors.APIError{
			Status:  http.StatusNotFound,
			Code:    api_errors.ErrCodeNotFound,
			Message: fmt.Sprintf(api_errors.ErrMsgNotFound, "android", id),
			Err:     err,
		}
	}

	return android, err
}

func (db DB) DeleteAndroid(tx *sql.Tx, ctx context.Context, id string) error {
	query := `DELETE FROM android WHERE id = $1 RETURNING id`
	var androidId string

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, id).Scan(&androidId)
	} else {
		err = db.Conn.QueryRowContext(ctx, query, id).Scan(&androidId)
	}

	if err == sql.ErrNoRows {
		return &api_errors.APIError{
			Status:  http.StatusNotFound,
			Code:    api_errors.ErrCodeNotFound,
			Message: fmt.Sprintf(api_errors.ErrMsgNotFound, "android", id),
			Err:     errors.New("no rows affected"),
		}
	}

	return err
}
