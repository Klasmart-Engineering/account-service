package main

import (
	"fmt"
	"kidsloop/account-service/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting KidsLoop Account Service on http://localhost:8080")

	router := SetUpRouter()
	router.Run()
}

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", HealthCheck)

	r.GET("/accounts/:id", handler.GetAccount)
	r.PUT("/accounts", handler.CreateAccount)
	r.DELETE("/accounts/:id", handler.DeleteAccount)
	return r
}

func HealthCheck(c *gin.Context) {
	c.String(200, "Server is running")
}
