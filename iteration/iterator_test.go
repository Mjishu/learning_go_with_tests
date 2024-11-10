package iterator

import "testing"

func TestIterator(t *testing.T) {
	t.Run("iterates as expected", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("runs for expected 3 times", func(t *testing.T) {
		repeated := Repeat("b", 3)
		expected := "bbb"

		assertCorrectMessage(t, repeated, expected)
	})
}

func assertCorrectMessage(t *testing.T, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("got %q but wanted %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
