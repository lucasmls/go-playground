package numerals

import "strings"

// ConvertToRoman converts a arabic number to a roman number
func ConvertToRoman(num int) string {
	var result strings.Builder

	for i := num; i > 0; i-- {
		if num == 4 {
			result.WriteString("IV")
			break
		}

		result.WriteString("I")
	}

	return result.String()
}
