package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	t.Run("returns a friendly message", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		RootHandler(response, request)

		got := response.Body.String()
		want := "Hello from account service"

		if got != "Hello from account service" {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
