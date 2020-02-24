package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{
		"name": "Laisla Pinto Coelho",
	}

	got := Search(dictionary, "name")
	want := "Laisla Pinto Coelho"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
