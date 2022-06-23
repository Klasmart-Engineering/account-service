package handler

import (
	"database/sql"
	"fmt"
	db "kidsloop/account-service/database"
	api_errors "kidsloop/account-service/errors"
	"kidsloop/account-service/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAndroid ... Create Android
// @Summary  Create a new account
// @Success  200  {object}  model.Android
// @Failure  404,500  {object}  model.ErrorResponse
// @Router   /accounts/{accountId}/android_groups/{groupId} [post]
func CreateAndroid(c *gin.Context) {
	type Uri struct {
		AccountID      string `uri:"accountId" binding:"required,uuid"`
		AndroidGroupID string `uri:"groupId" binding:"required,uuid"`
	}
	var uri Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.Error(err)
		return
	}

	response, err := db.RunInTransaction(c, func(tx *sql.Tx) (model.Android, error) {
		var android model.Android
		// validate that the account exists
		_, err := db.Database.GetAccount(nil, uri.AccountID)
		if err != nil {
			return android, err
		}

		// validate the android group exists for the specified account
		androidGroup, err := db.Database.GetAndroidGroup(nil, uri.AndroidGroupID)
		if err != nil {
			return android, err
		} else if androidGroup.AccountID != uri.AccountID {
			return android, &api_errors.APIError{
				Status:  http.StatusNotFound,
				Code:    api_errors.ErrCodeNotFound,
				Message: fmt.Sprintf(api_errors.ErrMsgNotFound, "android_group", uri.AndroidGroupID),
				Err:     err,
			}
		}

		return db.Database.CreateAndroid(nil, uri.AccountID, uri.AndroidGroupID)
	})

	if err != nil {
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, response)
	}
}
