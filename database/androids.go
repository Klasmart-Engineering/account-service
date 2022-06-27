package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	api_errors "kidsloop/account-service/errors"
	"kidsloop/account-service/model"
	"kidsloop/account-service/util"
	"net/http"

	"github.com/lib/pq"
)

func (db DB) CreateAndroid(tx *sql.Tx, ctx context.Context, androidGroupId string) (model.Android, error) {
	query := `INSERT INTO android (android_group_id) VALUES ($1) RETURNING id, android_group_id`
	android := model.Android{}

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, androidGroupId).Scan(&android.ID, &android.AndroidGroupID)
	} else {
		err = db.Conn.QueryRowContext(ctx, query, androidGroupId).Scan(&android.ID, &android.AndroidGroupID)
	}
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			// check for android group not found
			if err.Code == "23503" {
				// foreign key constraint violation: https://www.postgresql.org/docs/9.3/errcodes-appendix.html
				return android, &api_errors.APIError{
					Status:  http.StatusNotFound,
					Code:    api_errors.ErrCodeNotFound,
					Message: fmt.Sprintf(api_errors.ErrMsgNotFound, "android_group", androidGroupId),
					Err:     err,
				}
			}
		}
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

func (db DB) GetAndroidsByGroup(tx *sql.Tx, groupId string, offset int, pageSize int) ([]model.Android, error) {
	query := `SELECT id, android_group_id FROM android WHERE android_group_id = $1 ORDER BY "id" OFFSET $2 LIMIT $3`
	androids := []model.Android{}

	limit := pageSize
	if limit == 0 {
		limit = util.DefaultPageSize
	}

	rows, err := db.Conn.Query(query, groupId, offset, limit)
	if err != nil {
		return androids, err
	}
	for rows.Next() {
		var android model.Android
		err := rows.Scan(&android.ID, &android.AndroidGroupID)
		if err != nil {
			return androids, err
		}
		androids = append(androids, android)
	}

	return androids, err
}
