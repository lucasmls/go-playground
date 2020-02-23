package repeat

// Repeat ...
func Repeat(character string) string {
	repeatedCharacters := ""

	for i := 0; i < 5; i++ {
		repeatedCharacters += character
	}

	return repeatedCharacters
}
