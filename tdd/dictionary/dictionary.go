package dictionary

import "errors"

// Dictionary ...
type Dictionary map[string]string

// ErrNotFound ...
var ErrNotFound = errors.New("could not find the word you were looking for")

// Search ...
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
