package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRootHandler(t *testing.T) {
	t.Run("returns a friendly message", func(t *testing.T) {
		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		HealthCheck(c)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "account-service is running", response.Body.String())
	})
}
