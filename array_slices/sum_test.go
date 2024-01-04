package main

import (
	"reflect"
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
func TestSumAll(t *testing.T) {
	t.Run("Sum of all slices", func(t *testing.T) {
		slice1 := []int{6, 2, 1}
		slice2 := []int{36, 56, 5}
		slice3 := []int{1, 25, 511}

		got := SumAll(slice1, slice2, slice3)
		want := []int{Sum(slice1), Sum(slice2), Sum(slice3)}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestSumAll failed: got %v, want %v, given %v", got, want, []interface{}{slice1, slice2, slice3})
		}
	})
}
