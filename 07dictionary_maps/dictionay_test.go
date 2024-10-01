package main

import (
	"errors"
	"testing"
)

// create new types from existing ones
type Dictionary map[string]string

var TestValue string = "this is for test"
var ErrorNotFound = errors.New("error!, key not found")
var ErrorAlreadyExists = errors.New("error! key already exists")

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if ok {
		return value, nil
	}
	return "", ErrorNotFound
}

func (d Dictionary) Add(key string, value string) error {
	// d[key] = value

	_, err := d.Search(key)
	switch err {
	case ErrorNotFound:
		d[key] = value
	case nil:
		return ErrorAlreadyExists
	default:
		return nil
	}

	return nil
}

func AssertString(t testing.TB, got, want, searchKey string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q and want %q for given %q \n", got, want, searchKey)
	}
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func AssertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("got error %q, not expecting", got)
	}
}

func AssertAddedValue(t testing.TB, dictionary Dictionary, want, key string) {
	t.Helper()
	got, err := dictionary.Search(key)
	AssertNoError(t, err)
	AssertString(t, got, want, key)
}

func TestSearch(t *testing.T) {

	t.Run("known key", func(t *testing.T) {
		dictionary := Dictionary{"test": TestValue}
		searchKey := "test"
		got, err := dictionary.Search(searchKey)
		want := TestValue
		AssertNoError(t, err)
		AssertString(t, got, want, searchKey)
	})

	t.Run("unknown key", func(t *testing.T) {
		dictionary := Dictionary{"test": TestValue}
		searchKey := "xyz"
		_, err := dictionary.Search(searchKey)
		AssertError(t, err, ErrorNotFound)
	})

	t.Run("Add new key, value", func(t *testing.T) {
		dictionary := Dictionary{}
		searchKey := "xyz"
		err := dictionary.Add(searchKey, TestValue)
		AssertError(t, err, nil)
		AssertAddedValue(t, dictionary, TestValue, searchKey)
	})

	t.Run("Add existing key, value", func(t *testing.T) {
		dictionary := Dictionary{"test": TestValue}
		searchKey := "test"
		err := dictionary.Add(searchKey, TestValue)
		AssertError(t, err, ErrorAlreadyExists)
		AssertAddedValue(t, dictionary, TestValue, searchKey)
	})
}
