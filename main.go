package main

import (
	"fmt"
	_ "kidsloop/account-service/docs"
	"kidsloop/account-service/handler"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"   // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
)

// @title    accounts-service documentation
// @version  0.0.1
// @host     localhost:8080

func main() {
	fmt.Println("Starting account-service on http://localhost:8080")

	router := SetUpRouter()
	router.Run()
}

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", HealthCheck)

	r.GET("/accounts/:accountId", handler.GetAccount)
	r.PUT("/accounts", handler.CreateAccount)
	r.DELETE("/accounts/:accountId", handler.DeleteAccount)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

func HealthCheck(c *gin.Context) {
	c.String(200, "account-service is running")
}
