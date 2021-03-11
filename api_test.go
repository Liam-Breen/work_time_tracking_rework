package main

import "testing"

func TestAPI(t *testing.T) {

	t.Run("Starting server message", func(t *testing.T) {
		got := start_message()
		want := "STARTING SERVER..."

		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	})

}
