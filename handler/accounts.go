package handler

import (
	"database/sql"
	db "kidsloop/account-service/database"
	"kidsloop/account-service/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAccount ... Create Account
// @Summary  Create a new account and associated android group and android. All newly created entities are returned.
// @Tags     accounts
// @Success  200  {object}  model.CreateAccountResponse
// @Failure  500  {object}  model.ErrorResponse
// @Router   /accounts [post]
func CreateAccount(c *gin.Context) {
	response, err := db.RunInTransaction(c, func(tx *sql.Tx) (*model.CreateAccountResponse, error) {
		account, err := db.Database.CreateAccount(tx)
		if err != nil {
			return nil, err
		}

		androidGroup, err := db.Database.CreateAndroidGroup(tx, account.ID)
		if err != nil {
			return nil, err
		}

		android, err := db.Database.CreateAndroid(tx, androidGroup.ID)
		if err != nil {
			return nil, err
		}

		return &model.CreateAccountResponse{
			Account:      account,
			Android:      android,
			AndroidGroup: androidGroup,
		}, nil
	})

	if err != nil {
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// GetAccount ... Get Account
// @Summary  Get details of an account
// @Tags     accounts
// @Param    id           path      string  true  "Account ID"
// @Success  200          {object}  model.Account
// @Failure  400,404,500  {object}  model.ErrorResponse
// @Router   /accounts/{id} [get]
func GetAccount(c *gin.Context) {
	type Uri struct {
		ID string `uri:"id" binding:"required,uuid"`
	}
	var uri Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.Error(err)
		return
	}

	account, err := db.Database.GetAccount(nil, uri.ID)

	if err != nil {
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, account)
	}
}

// DeleteAccount ... Delete Account
// @Summary  Deletes an account and all associated android groups and androids
// @Tags     accounts
// @Param    id           path      string  true  "Account ID"
// @Success  200          {object}  model.Account
// @Failure  400,404,500  {object}  model.ErrorResponse
// @Router   /accounts/{id} [delete]
func DeleteAccount(c *gin.Context) {
	type Uri struct {
		ID string `uri:"id" binding:"required,uuid"`
	}
	var uri Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.Error(err)
		return
	}

	account, err := db.Database.DeleteAccount(nil, uri.ID)
	if err != nil {
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, account)
	}
}
