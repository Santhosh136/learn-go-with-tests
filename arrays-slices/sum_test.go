package arraysslices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("sum of numbers slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	t.Run("sum of all slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{2, 5}, []int{4, 0, 2})
		want := []int{6, 7, 6}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSlicesAreEqual := func(t *testing.T, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum of all slices tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4}, []int{1, 3, 4})
		want := []int{9, 7}
		checkSlicesAreEqual(t, got, want)
	})

	t.Run("with some empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{})
		want := []int{5, 0}
		checkSlicesAreEqual(t, got, want)
	})
}
