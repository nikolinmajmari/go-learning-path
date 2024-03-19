package dictionary

import (
	"errors"
)

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you're looking for ")
	ErrWordExists       = DictionaryErr("could not add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

/*
var (

	ErrNotFound   = errors.New("could not find item you're searching for")
	ErrWordExists = errors.New("can not add word because it already exists")

)
*/
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch {
	case errors.Is(err, ErrNotFound):
		d[word] = definition
	case err == nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch {
	case errors.Is(err, ErrNotFound):
		return ErrWordDoesNotExist
	case err == nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func Search(dict map[string]string, word string) string {
	return dict[word]
}
