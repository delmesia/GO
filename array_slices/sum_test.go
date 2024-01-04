package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("TestSum failed: got %d, want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Sum of slices excluding head", func(t *testing.T) {
		slice1 := []int{1, 5, 5, 5, 5}
		slice2 := []int{1, 5, 5, 5}

		got := SumAllTails(slice1, slice2)
		want := []int{20, 15}

		checkSums(t, got, want)
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{0, 1})
		want := []int{0, 1}

		checkSums(t, got, want)
	})

}
