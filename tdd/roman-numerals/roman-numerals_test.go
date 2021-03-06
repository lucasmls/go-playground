package numerals

import (
	"fmt"
	"testing"
)

var (
	cases = []struct {
		Arabic int
		Roman  string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{11, "XI"},
		{12, "XII"},
		{13, "XIII"},
		{14, "XIV"},
		{18, "XVIII"},
		{20, "XX"},
		{39, "XXXIX"},
		{50, "L"},
		{100, "C"},
		{500, "D"},
		{1000, "M"},
		{1984, "MCMLXXXIV"},
		{798, "DCCXCVIII"},
		{1006, "MVI"},
	}
)

func Test_ConvertToRoman(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("convert %d to %s", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			want := test.Roman

			if got != want {
				t.Errorf(`got %q want %q`, got, want)
			}
		})
	}
}
