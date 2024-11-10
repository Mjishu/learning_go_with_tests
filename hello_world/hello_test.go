package hello_world

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Mjishu", "")
		want := "Hello, Mjishu"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello with an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("test runs in spanish!", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})
	t.Run("test runs in french!", func(t *testing.T) {
		got := Hello("meowe", "French")
		want := "Bonjour, meowe"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("go %q but wanted %q", got, want)
	}
}
