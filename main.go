package main

import (
	"fmt"
	"kidsloop/account-service/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting KidsLoop Account Service")

	r := gin.Default()
	r.GET("/", HealthCheck)

	r.GET("/accounts/:accountId", handler.GetAccount)
	r.PUT("/accounts/", handler.CreateAccount)
	r.DELETE("/accounts/:accountId", handler.DeleteAccount)

	r.Run()
}

func HealthCheck(c *gin.Context) {
	c.String(200, "Server is running")
}
