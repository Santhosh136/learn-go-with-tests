package maps

import "testing"

func TestSearch(t *testing.T) {

	t.Run("search dictionary for word", func(t *testing.T) {
		dictionary := Dictionary{"test": "testing word"}

		meaning, _ := dictionary.Search("test")
		want := "testing word"

		assertStrings(t, meaning, want)
	})

	t.Run("search dictionary unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "testing word"}
		_, err := dictionary.Search("unknown")
		assertErrors(t, err, ErrorNotFound)
	})

	t.Run("add new word to dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "testing word"}
		err := dictionary.AddNewWord("new", "new word")

		got := len(dictionary)
		want := 2

		assertIntegers(t, got, want)
		assertErrors(t, err, nil)
	})

	t.Run("add existing word to dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "testing word"}
		err := dictionary.AddNewWord("test", "new testing word")

		got := len(dictionary)
		want := 1

		if err == nil {
			t.Fatal("expected error but didn't get one")
		}

		assertIntegers(t, got, want)
		assertErrors(t, err, ErrorAlreadyExists)
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s, given %s", got, want, "test")
	}
}

func assertIntegers(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q but want %q", got, want)
	}
}
