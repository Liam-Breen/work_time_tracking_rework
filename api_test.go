package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI(t *testing.T) {

	t.Run("Starting server message", func(t *testing.T) {
		got := startMessage()
		want := "STARTING SERVER..."

		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	})

	t.Run("Return Liam Breen's hours worked", func(t *testing.T) {
		request := newGetUserHours("LiamBreen")
		response := httptest.NewRecorder()

		UserServer(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("Return James Bos' hours worked", func(t *testing.T) {
		request := newGetUserHours("JamesBos")
		response := httptest.NewRecorder()

		UserServer(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

}

func newGetUserHours(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/user/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
