package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	fmt.Println("CreateAccount")
	c.JSON(200, gin.H{
		"accountId": "abc123",
	})
}

func GetAccount(c *gin.Context) {
	fmt.Println("GetAccount")
	accountId := c.Param("accountId")
	c.JSON(200, gin.H{
		"accountId": accountId,
	})
}

func DeleteAccount(c *gin.Context) {
	fmt.Println("DeleteAccount")
	accountId := c.Param("accountId")
	c.JSON(200, gin.H{
		"accountId": accountId,
	})
}
