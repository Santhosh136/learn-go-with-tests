package maps

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrWordNotFound    = DictionaryErr("could not find the word")
	ErrorAlreadyExists = DictionaryErr("word already exists")
	ErrorDoesNotExists = DictionaryErr("word does not exists")
)

func (d Dictionary) Search(word string) (string, error) {
	meaning, found := d[word]
	if !found {
		return "", ErrWordNotFound
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

func (d Dictionary) UpdateWord(word, newMeaning string) error {
	_, err := d.Search(word)

	if err == ErrWordNotFound {
		return ErrorDoesNotExists
	}

	d[word] = newMeaning
	return nil
}

func (d Dictionary) DeleteWord(word string) error {
	_, err := d.Search(word)

	if err == ErrWordNotFound {
		return ErrorDoesNotExists
	}

	delete(d, word)
	return nil
}
