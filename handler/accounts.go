package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	fmt.Println("CreateAccount")
	c.JSON(200, gin.H{
		"account_id": "abc123",
	})
}

func GetAccount(c *gin.Context) {
	fmt.Println("GetAccount")
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id": id,
	})
}

func DeleteAccount(c *gin.Context) {
	fmt.Println("DeleteAccount")
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id": id,
	})
}
