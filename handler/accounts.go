package handler

import (
	"fmt"
	"kidsloop/account-service/model"

	"github.com/gin-gonic/gin"
)

// CreateAccount ... Create Account
// @Summary  Create a new account
// @Success  200      {object}  model.Account
// @Failure  400,500  {object}  model.APIError
// @Router   /accounts [put]
func CreateAccount(c *gin.Context) {
	fmt.Println("CreateAccount")
	c.JSON(200, gin.H{
		"accountId": "abc123",
	})
}

// GetAccount ... Get Account
// @Summary  Get details of an account
// @Param    accountId  path      string  true  "Account ID"
// @Success  200        {object}  model.Account
// @Failure  400,500    {object}  model.APIError
// @Router   /accounts/{accountId} [get]
func GetAccount(c *gin.Context) {
	fmt.Println("GetAccount")
	accountId := c.Param("accountId")
	account := model.Account{
		ID: accountId,
	}
	c.JSON(200, account)
}

// DeleteAccount ... Delete Account
// @Summary  Delete an account
// @Param    accountId  query     string  true  "Account ID"
// @Success  200        {object}  model.Account
// @Failure  400,500    {object}  model.APIError
// @Router   /accounts/{accountId} [delete]
func DeleteAccount(c *gin.Context) {
	fmt.Println("DeleteAccount")
	accountId := c.Param("accountId")
	c.JSON(200, gin.H{
		"accountId": accountId,
	})
}
