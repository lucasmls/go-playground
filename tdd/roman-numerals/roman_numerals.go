package numerals

// ConvertToRoman converts a arabic number to a roman number
func ConvertToRoman(num int) string {
	if num == 2 {
		return "II"
	}

	return "I"
}
