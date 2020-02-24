package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"name": "Laisla Pinto Coelho",
	}

	got := dictionary.Search(dictionary, "name")
	want := "Laisla Pinto Coelho"

	assertString(t, got, want)
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
