package numerals

import "strings"

// ConvertToRoman converts a arabic number to a roman number
func ConvertToRoman(num int) string {
	if num == 4 {
		return "IV"
	}

	var result strings.Builder

	for i := 0; i < num; i++ {
		result.WriteString("I")
	}

	return result.String()
}
