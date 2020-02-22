package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to the correct person", func(t *testing.T) {
		got := Hello("Lucas")
		want := "Hello, Lucas"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello, World when a person is not supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

}
