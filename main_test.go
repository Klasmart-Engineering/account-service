package main

import (
	"net/http"
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
		want := "account-service is running"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

type Route struct {
	path   string
	method string
}

func TestExpectedRoutes(t *testing.T) {
	expectedRoutes := []Route{
		{
			path:   "/",
			method: "GET",
		},
		{
			path:   "/accounts/abc123",
			method: "GET",
		},
		{
			path:   "/accounts",
			method: "PUT",
		},
		{
			path:   "/accounts/abc123",
			method: "DELETE",
		},
	}

	router := SetUpRouter()

	for _, route := range expectedRoutes {
		req, _ := http.NewRequest(route.method, route.path, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		got := w.Code
		want := http.StatusOK

		if got != want {
			t.Errorf("got %d, want %d for route %s", got, want, route.path)
		}
	}
}

func TestUnexpectedRoutes(t *testing.T) {
	expectedRoutes := []Route{
		{
			path:   "/",
			method: "POST",
		},
		{
			path:   "/hello",
			method: "GET",
		},
		{
			path:   "/accounts",
			method: "DELETE",
		},
	}

	router := SetUpRouter()

	for _, route := range expectedRoutes {
		req, _ := http.NewRequest(route.method, route.path, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		got := w.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}
