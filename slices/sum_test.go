package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Test with input of any size", func(t *testing.T) {
		//Slice of ints
		slice := []int{1, 2, 3}
		got := Sum(slice)
		want := 6
		if got != want {
			t.Errorf("got %v, want %v, given%v", got, want, slice)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
