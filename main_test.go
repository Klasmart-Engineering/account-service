package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRootHandler(t *testing.T) {
	t.Run("returns a friendly message", func(t *testing.T) {
		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		HealthCheck(c)

		got := response.Body.String()
		want := "Server is running"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
