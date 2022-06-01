package main

import (
	"fmt"
	"kidsloop/account-service/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting KidsLoop Account Service")

	r := gin.Default()
	r.GET("/", HealthCheck)

	r.GET("/accounts/:accountId", controllers.GetAccount)
	r.PUT("/accounts/", controllers.CreateAccount)
	r.DELETE("/accounts/:accountId", controllers.DeleteAccount)

	r.Run()
}

func HealthCheck(c *gin.Context) {
	c.String(200, "Server is running")
}
