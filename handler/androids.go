package handler

import (
	"context"
	db "kidsloop/account-service/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// GetAndroid ... Get Android
// @Summary  Get details of an android
// @Tags     androids
// @Param    id           path      string  true  "Android ID"
// @Success  200          {object}  model.Android
// @Failure  400,404,500  {object}  model.ErrorResponse
// @Router   /androids/{id} [get]
func GetAndroid(c *gin.Context) {
	type Uri struct {
		ID string `uri:"id" binding:"required,uuid"`
	}
	var uri Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.Error(err)
		return
	}

	nrTxn := nrgin.Transaction(c)
	nrCtx := newrelic.NewContext(context.Background(), nrTxn)

	id := c.Param("id")
	android, err := db.Database.GetAndroid(nil, nrCtx, id)

	if err != nil {
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, android)
	}
}
