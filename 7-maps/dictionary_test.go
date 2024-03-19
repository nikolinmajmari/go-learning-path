package dictionary

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test value"}
	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is a test value"
		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "number"
		definition := "Definition of number"
		err := dictionary.Add(word, definition)
		if err != nil {
			t.Fatal("An unexpected error occured")
		}
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "existing"
		definition := "existing new definition"
		dictionary := Dictionary{
			word: definition,
		}
		err := dictionary.Add(word, definition)
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	word := "word"
	definition := "Definition of word"
	dictionary := Dictionary{word: definition}

	newDefinition := "New definition"

	err := dictionary.Update(word, newDefinition)
	if err != nil {
		t.Fatal("An unexpected error occurred")
	}
	assertDefinition(t, dictionary, word, newDefinition)
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("Should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertNoError(t testing.TB, err error) {
	if err != nil {
		t.Errorf("got an unexpected error %s", err.Error())
	}
}
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
