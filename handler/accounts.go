package handler

import (
	db "kidsloop/account-service/database"

	"github.com/gin-gonic/gin"
)

// CreateAccount ... Create Account
// @Summary  Create a new account
// @Success  200      {object}  model.Account
// @Failure  400,500  {object}  model.APIError
// @Router   /accounts [put]
func CreateAccount(c *gin.Context) {
	account, err := db.Database.CreateAccount()
	if err != nil {
		c.String(500, err.Error())
	}
	c.JSON(200, account)
}

// GetAccount ... Get Account
// @Summary  Get details of an account
// @Param    id  path      string  true  "Account ID"
// @Success  200        {object}  model.Account
// @Failure  400,500    {object}  model.APIError
// @Router   /accounts/{id} [get]
func GetAccount(c *gin.Context) {
	id := c.Param("id")
	account, err := db.Database.GetAccount(id)
	if err != nil {
		c.String(500, err.Error())
	}
	c.JSON(200, account)
}

// DeleteAccount ... Delete Account
// @Summary  Delete an account
// @Param    id  query     string  true  "Account ID"
// @Success  200        {object}  model.Account
// @Failure  400,500    {object}  model.APIError
// @Router   /accounts/{id} [delete]
func DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	err := db.Database.DeleteAccount(id)
	if err != nil {
		c.String(500, err.Error())
	}
	c.String(200, "Success")
}
