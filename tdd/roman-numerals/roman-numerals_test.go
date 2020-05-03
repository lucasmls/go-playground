package numerals

import "testing"

func TestRomanNumerals(t *testing.T) {

	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{`1 gets converted to I`, 1, "I"},
		{`2 gets converted to II`, 2, "II"},
		{`3 gets converted to III`, 3, "III"},
		{`4 gets converted to IV (can't repeat more than 3 times)`, 4, "IV"},
		{`5 gets converted to V`, 5, "V"},
		{`6 gets converted to VI`, 6, "VI"},
		{`7 gets converted to VII`, 7, "VII"},
		{`8 gets converted to VIII`, 8, "VIII"},
		{`9 gets converted to IX`, 9, "IX"},
		{`10 gets converted to X`, 10, "X"},
		{`11 gets converted to XI`, 11, "XI"},
		{`12 gets converted to XII`, 12, "XII"},
		{`13 gets converted to XIII`, 13, "XIII"},
		{`14 gets converted to XIV`, 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{"50 gets converted to L", 50, "L"},
		{"100 gets converted to C", 100, "C"},
		{"500 gets converted to D", 500, "D"},
		{"1000 gets converted to M", 1000, "M"},
		{"1984 gets converted to MCMLXXXIV", 1984, "MCMLXXXIV"},
		{"798 gets converted to DCCXCVIII", 798, "DCCXCVIII"},
		{"1006 gets converted to MVI", 1006, "MVI"},
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
