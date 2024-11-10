package integers

import "testing"

func TestAdding(t *testing.T) {
	t.Run("Adds the 2 inputs correctly", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertCorrectMessage(t, sum, expected)
	})
}

func assertCorrectMessage(t *testing.T, sum, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("Expected %d but got %d!", expected, sum)
	}
}
