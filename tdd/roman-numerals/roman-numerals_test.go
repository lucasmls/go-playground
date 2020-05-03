package numerals

import "testing"

func TestRomanNumerals(t *testing.T) {

	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{
			Description: `1 gets converted to I`,
			Arabic:      1,
			Want:        "I",
		},
		{
			Description: `2 gets converted to II`,
			Arabic:      2,
			Want:        "II",
		},
		{
			Description: `3 gets converted to III`,
			Arabic:      3,
			Want:        "III",
		},
		{
			Description: `4 gets converted to IV (can't repeat more than 3 times)`,
			Arabic:      4,
			Want:        "IV",
		},
		{
			Description: `5 gets converted to V`,
			Arabic:      5,
			Want:        "V",
		},
		{
			Description: `6 gets converted to VI`,
			Arabic:      6,
			Want:        "VI",
		},
		{
			Description: `7 gets converted to VII`,
			Arabic:      7,
			Want:        "VII",
		},
		{
			Description: `8 gets converted to VIII`,
			Arabic:      8,
			Want:        "VIII",
		},
		{
			Description: `9 gets converted to IX`,
			Arabic:      9,
			Want:        "IX",
		},
		{
			Description: `10 gets converted to X`,
			Arabic:      10,
			Want:        "X",
		},
		{
			Description: `11 gets converted to XI`,
			Arabic:      11,
			Want:        "XI",
		},
		{
			Description: `12 gets converted to XII`,
			Arabic:      12,
			Want:        "XII",
		},
		{
			Description: `13 gets converted to XIII`,
			Arabic:      13,
			Want:        "XIII",
		},
		{
			Description: `14 gets converted to XIV`,
			Arabic:      14,
			Want:        "XIV",
		},
		{
			Description: "18 gets converted to XVIII",
			Arabic:      18,
			Want:        "XVIII",
		},
		{
			Description: "20 gets converted to XX",
			Arabic:      20,
			Want:        "XX",
		},
		{
			Description: "39 gets converted to XXXIX",
			Arabic:      39,
			Want:        "XXXIX",
		},
		{
			Description: "50 gets converted to L",
			Arabic:      50,
			Want:        "L",
		},
		{
			Description: "100 gets converted to C",
			Arabic:      100,
			Want:        "C",
		},
		{
			Description: "500 gets converted to D",
			Arabic:      500,
			Want:        "D",
		},
		{
			Description: "1000 gets converted to M",
			Arabic:      1000,
			Want:        "M",
		},
		{
			Description: "1984 gets converted to MCMLXXXIV",
			Arabic:      1984,
			Want:        "MCMLXXXIV",
		},
		{
			Description: "798 gets converted to DCCXCVIII",
			Arabic:      798,
			Want:        "DCCXCVIII",
		},
		{
			Description: "1006 gets converted to MVI",
			Arabic:      1006,
			Want:        "MVI",
		},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			want := test.Want

			if got != want {
				t.Errorf(`got %q want %q`, got, want)
			}
		})
	}

}
