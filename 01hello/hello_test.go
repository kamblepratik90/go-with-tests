package main

import "testing"

// func TestHello(t *testing.T) {
// 	got := Hello()
// 	want := "Hello World"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// func TestHello2(t *testing.T) {
// 	got := Hello("My Friends")
// 	want := "Hello, My Friends"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

func TestHello3(t *testing.T) {
	t.Run("test simple hello with msg", func(t *testing.T) {

		got := Hello("Batman")
		want := "Hello, Batman"

		// if got != want {
		// 	t.Errorf("got %q and want %q", got, want)
		// }
		assertCorrectMessage(t, got, want)
	})

	t.Run("test hello world without msg", func(t *testing.T) {

		got := Hello("")
		want := "Hello, World"

		// if got != want {
		// 	t.Errorf("got %q and want %q", got, want)
		// }
		assertCorrectMessage(t, got, want)
	})
}

// func assertCorrectMessage(t testing.TB, got string, want string) {
func assertCorrectMessage(t testing.TB, got, want string) {

	t.Helper() // to tell the test suite that this method is a helper
	// By doing this, when it fails, the line number reported will be in our function call
	// rather than inside our test helper
	if got != want {
		t.Errorf("got %q and want %q", got, want)
	}
}
