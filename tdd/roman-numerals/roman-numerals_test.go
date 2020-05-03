package numerals

import "testing"

func TestRomanNumerals(t *testing.T) {

	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{
			Description: `"1" gets converted to "I"`,
			Arabic:      1,
			Want:        "I",
		},
		{
			Description: `"2" gets converted to "II"`,
			Arabic:      2,
			Want:        "II",
		},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			want := test.Want

			if got != want {
				t.Errorf(`got "%q" want "%q"`, got, want)
			}
		})
	}
}
