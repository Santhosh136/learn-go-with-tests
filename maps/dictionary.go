package maps

import "errors"

type Dictionary map[string]string

var ErrorNotFound = errors.New("could not find the word")
var ErrorAlreadyExists = errors.New("word already exists")

func (d Dictionary) Search(word string) (string, error) {
	meaning, found := d[word]
	if !found {
		return "", ErrorNotFound
	}
	return meaning, nil
}

func (d Dictionary) AddNewWord(word, meaning string) error {
	_, err := d.Search(word)

	if err == nil {
		return ErrorAlreadyExists
	}

	d[word] = meaning
	return nil
}
