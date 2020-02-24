package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"laisla": "Laisla Pinto Coelho"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("laisla")
		want := "Laisla Pinto Coelho"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("lucas", "Lucas Mendes")

	want := "Lucas Mendes"
	got, err := dictionary.Search("lucas")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertString(t, got, want)
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
