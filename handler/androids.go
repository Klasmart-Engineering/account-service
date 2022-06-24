package handler

import (
	db "kidsloop/account-service/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAndroid ... Create Android
// @Summary  Create a new account
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

	android, err := db.Database.CreateAndroid(nil, uri.ID)

	if err != nil {
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, android)
	}
}
