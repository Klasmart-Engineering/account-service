package handler

import (
	"database/sql"
	db "kidsloop/account-service/database"
	"kidsloop/account-service/model"

	"github.com/gin-gonic/gin"
)

// CreateAccount ... Create Account
// @Summary  Create a new account
// @Success  200      {object}  model.CreateAccountResponse
// @Failure  400,500  {object}  model.APIError
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

		android, err := db.Database.CreateAndroid(tx, account.ID, androidGroup.ID)
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
		c.String(500, err.Error())
	} else {
		c.JSON(200, response)
	}
}

// GetAccount ... Get Account
// @Summary  Get details of an account
// @Param    id  path      string  true  "Account ID"
// @Success  200        {object}  model.Account
// @Failure  400,500    {object}  model.APIError
// @Router   /accounts/{id} [get]
func GetAccount(c *gin.Context) {
	id := c.Param("id")
	account, err := db.Database.GetAccount(nil, id)
	if err != nil {
		c.String(500, err.Error())
	} else {
		c.JSON(200, account)
	}
}

// DeleteAccount ... Delete Account
// @Summary  Delete an account
// @Param    id  path     string  true  "Account ID"
// @Success  200        {object}  model.Account
// @Failure  400,500    {object}  model.APIError
// @Router   /accounts/{id} [delete]
func DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	err := db.Database.DeleteAccount(nil, id)
	if err != nil {
		c.String(500, err.Error())
	} else {
		c.String(200, "Success")

	}
}
