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
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertString(t, err.Error(), want)
	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
