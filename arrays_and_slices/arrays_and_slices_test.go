package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	checksums := func(t testing.TB, expected, res []int) {
		t.Helper()
		if !reflect.DeepEqual(expected, res) {
			t.Errorf("Expected %v result %v", expected, res)
		}
	}

	t.Run("Sum of any size", func(t *testing.T) {
		number := []int{1, 2, 3}
		expected := 6
		result := Sum(number)
		if result != expected {
			t.Errorf("Expected %d and result %d", expected, result)
		}
	})

	t.Run("Sum of all slices", func(t *testing.T) {
		expected := []int{6, 9}
		res := SumAll([]int{1, 2, 3}, []int{5, 4})
		checksums(t, expected, res)
	})

	t.Run("Sum all trails", func(t *testing.T) {
		expected := SumAllTrails([]int{1, 2}, []int{0, 9})
		res := []int{2, 9}
		checksums(t, expected, res)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTrails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checksums(t, got, want)
	})
}
