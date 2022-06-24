package db

import (
	"context"
	"database/sql"
	"kidsloop/account-service/model"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func (db DB) CreateAndroidGroup(tx *sql.Tx, accountID string, nrTxn *newrelic.Transaction) (model.AndroidGroup, error) {
	query := `INSERT INTO android_group (account_id) VALUES ($1) RETURNING id`
	androidGroup := model.AndroidGroup{}

	nrCtx := newrelic.NewContext(context.Background(), nrTxn)

	var err error
	if tx != nil {
		err = tx.QueryRowContext(nrCtx, query, accountID).Scan(&androidGroup.ID)
	} else {
		err = db.Conn.QueryRowContext(nrCtx, query, accountID).Scan(&androidGroup.ID)
	}

	return androidGroup, err
}
