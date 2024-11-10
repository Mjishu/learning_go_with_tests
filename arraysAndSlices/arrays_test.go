package arraySlices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of size 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectMessage(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Runs test as expected", func(t *testing.T) {
		got := SumAll([]int{3, 2}, []int{0, 9})
		want := []int{5, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v but wanted %v", got, want)
		}
	})
}

func assertCorrectMessage(t *testing.T, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %d but wanted %d given %v", got, want, numbers)
	}
}
