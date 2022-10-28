package mydict

import (
	"errors"
)

// Dictionary type
type Dictionary /** alias for type **/ map[string]string

var (
	errNotFound      = errors.New("not found")
	errAlreadyExists = errors.New("word already exists")
	errCantUpdate    = errors.New("cannot update not-existing word")
	errCantDelete    = errors.New("cannot delete not-existing word")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word string, definition string) error {

	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = definition
	case nil:
		return errAlreadyExists
	}
	return nil
}

// Update a word with given definition
func (d Dictionary) Update(word, newDefinition string) error {

	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = newDefinition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word from dictionary
func (d Dictionary) Delete(word string) error {

	_, err := d.Search(word)

	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil

}
