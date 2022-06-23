package handler

import (
	"fmt"
	"net/http"

	api_errors "kidsloop/account-service/errors"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.Use(api_errors.ErrorHandler)

	// Recover from panics and send a 500 error
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

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
