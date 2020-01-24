package strings

import "testing"

import "strings"

const errorMessage = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	tests := []struct {
		text     string
		part     string
		expected int
	}{
		{"Lucas Mendes", "Lucas", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Laisla", "a", 1},
	}

	for _, test := range tests {
		actual := strings.Index(test.text, test.part)

		if actual != test.expected {
			t.Errorf(errorMessage, test.text, test.part, test.expected, actual)
		}
	}
}
