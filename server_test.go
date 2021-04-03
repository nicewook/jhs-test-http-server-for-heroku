package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreetingHandler(t *testing.T) {

	t.Run("testing hello", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Error(err)
		}
		resp := httptest.NewRecorder()

		GreetingHandler(resp, req)

		got := resp.Body.String()
		want := "hello"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
