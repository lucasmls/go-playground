package dictionary

// Dictionary ...
type Dictionary map[string]string

// Search ...
func (d Dictionary) Search(dictionary map[string]string, word string) string {
	return d[word]
}
