package main

import "testing"

// func TestHello(t *testing.T) {
// 	got := Hello()
// 	want := "Hello World"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

func TestHello2(t *testing.T) {
	got := Hello("My Friends")
	want := "Hello My Friends"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
