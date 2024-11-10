package arraySlices

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of size 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectMessage(t, got, want, numbers)
	})
}

func assertCorrectMessage(t *testing.T, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %d but wanted %d given %v", got, want, numbers)
	}
}
