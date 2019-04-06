package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		// By doing this when it fails the line number reported will be in our function call
		// rather than inside our test helper
		t.Helper()

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Tom")
		want := "Hello, Tom"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in spain". func(t, *testing.T)  {
		// prepare to write inter function
	})
}
