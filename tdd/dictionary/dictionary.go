package dictionary

// DictionaryErr ...
type DictionaryErr string

func (de DictionaryErr) Error() string {
	return string(de)
}

const (
	// ErrNotFound ...
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists ...
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	// ErrWordDoesNotExists ...
	ErrWordDoesNotExists = DictionaryErr("cannot update a word that are not registered")
)

// Dictionary ...
type Dictionary map[string]string

// Search ...
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add ...
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update ...
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}
