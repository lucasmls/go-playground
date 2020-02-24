package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"name": "Laisla Pinto Coelho"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("name")
		want := "Laisla Pinto Coelho"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
