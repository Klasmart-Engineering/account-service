package handler

import (
	"context"
	db "kidsloop/account-service/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// CreateAndroid ... Create Android
// @Summary  Create a new android in the specified android group
// @Tags     androids
// @Param    id       path      string  true  "Android Group ID"
// @Success  200      {object}  model.Android
// @Failure  404,500  {object}  model.ErrorResponse
// @Router   /android_groups/{id}/androids [post]
func CreateAndroid(c *gin.Context) {
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

	android, err := db.Database.CreateAndroid(nil, nrCtx, uri.ID)

	if err != nil {
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, android)
	}
}
