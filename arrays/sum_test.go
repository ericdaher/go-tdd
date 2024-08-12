package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	checkSum := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %v, wanted %v given %v", got, want, numbers)
		}
	}

	checkSum(t, got, want)
}

func TestSumAllUsingMake(t *testing.T) {
	slice_1 := []int{1, 2}
	slice_2 := []int{0, 9}

	got := SumAllUsingMake(slice_1, slice_2)
	want := []int{3, 9}

	assertSlicesEqual(t, got, want, slice_1, slice_2)
}

func TestSumAllUsingAppend(t *testing.T) {
	slice_1 := []int{1, 2}
	slice_2 := []int{0, 9}

	got := SumAllUsingAppend(slice_1, slice_2)
	want := []int{3, 9}

	assertSlicesEqual(t, got, want, slice_1, slice_2)
}

func TestSumAllTails(t *testing.T) {
	t.Run("when passing regular slices", func(t *testing.T) {
		slice_1 := []int{1, 2}
		slice_2 := []int{0, 9}

		got := SumAllTails(slice_1, slice_2)
		want := []int{2, 9}

		assertSlicesEqual(t, got, want, slice_1, slice_2)
	})

	t.Run("when passing one empty slice", func(t *testing.T) {
		slice_1 := []int{1, 2}
		slice_2 := []int{}

		got := SumAllTails(slice_1, slice_2)
		want := []int{2, 0}

		assertSlicesEqual(t, got, want, slice_1, slice_2)
	})
}

func TestSumAllHeads(t *testing.T) {
	slice_1 := []int{1, 2}
	slice_2 := []int{0, 9}

	got := SumAllHeads(slice_1, slice_2)
	want := []int{1, 0}

	assertSlicesEqual(t, got, want, slice_1, slice_2)
}

func assertSlicesEqual(t testing.TB, got, want, slice_1, slice_2 []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v, wanted %v given %v and %v", got, want, slice_1, slice_2)
	}
}
