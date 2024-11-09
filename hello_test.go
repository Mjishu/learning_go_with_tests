package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Mjishu")
		want := "Hello, Mjishu"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello with an empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("go %q but wanted %q", got, want)
	}
}
