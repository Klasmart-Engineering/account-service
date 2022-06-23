package handler

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", HealthCheck)

	r.GET("/accounts/:id", GetAccount)
	r.POST("/accounts", CreateAccount)
	r.DELETE("/accounts/:id", DeleteAccount)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

func HealthCheck(c *gin.Context) {
	c.String(200, "account-service is running")
}
