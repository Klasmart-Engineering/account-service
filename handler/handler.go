package handler

import (
	api_errors "kidsloop/account-service/errors"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.Use(api_errors.ErrorHandler)

	r.GET("/", HealthCheck)

	r.GET("/accounts/:id", GetAccount)
	r.POST("/accounts", CreateAccount)
	r.DELETE("/accounts/:id", DeleteAccount)

	r.POST("/android_groups/:id/androids", CreateAndroid)
	r.GET("/androids/:id", GetAndroid)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

func HealthCheck(c *gin.Context) {
	c.String(200, "account-service is running")
}
