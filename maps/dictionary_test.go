package main

import (
	"testing"
)

func TestSearch(t *testing.T) {

	dict := Dictionary{"Test": "This is just a test"}

	t.Run("known key", func(t *testing.T) {
		got, _ := dict.Search("Test")
		want := "This is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, got := dict.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new key", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("Key1", "Value1")
		assertError(t, err, nil)
		assertDefinition(t, dictionary, "Key1", "Value1")
	})
	t.Run("duplicate key", func(t *testing.T) {
		dictionary := Dictionary{"Key1": "Value1"}
		err := dictionary.Add("Key1", "Value1")
		assertError(t, err, ErrDupKey)
		assertDefinition(t, dictionary, "Key1", "Value1")
	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing key", func(t *testing.T) {
		dictionary := Dictionary{"Key1": "Value1"}
		err := dictionary.Update("Key1", "New Value")
		assertError(t, err, nil)
		assertDefinition(t, dictionary, "Key1", "New Value")
	})

	t.Run("new key", func(t *testing.T) {
		dictionary := Dictionary{"Key1": "Value1"}
		err := dictionary.Update("Key2", "Updated Value")
		assertError(t, err, ErrUpdateFailed)
	})

}

func TestDelete(t *testing.T) {
	t.Run("delete existing key", func(t *testing.T) {
		dictionary := Dictionary{"Key1": "Value1"}
		dictionary.Delete("Key1")
		_, err := dictionary.Search("Key1")
		if err == nil {
			t.Error("Key should already be deleted")
		}
	})

	t.Run("delete unknown key", func(t *testing.T) {
		dictionary := Dictionary{"Key1": "Value1"}
		err := dictionary.Delete("Key2")
		assertError(t, err, ErrDeleteFailed)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find the added word:", err)
	}
	assertString(t, got, definition)
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertString(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}
}
