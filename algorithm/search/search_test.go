package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		search, want int
	}{
		{-1, 0},
		{200, 10},
		{4, 7},
		{3, 5},
		{10, -1},
		{199, 9},
		{5, 8},
		{2, 3},
	}
	arry := []int{-1, 1, 1, 2, 2, 3, 3, 4, 5, 199, 200}
	for _, c := range cases {
		result := BinarySearch(arry, c.search)
		if result != c.want {
			t.Errorf("BinarySearch %d in %v, want %d, but return %d", c.search, arry, c.want, result)
		}
	}
}

func TestBinarySearchNest(t *testing.T) {
	cases := []struct {
		search, want int
	}{
		{-1, 0},
		{200, 10},
		{4, 7},
		{3, 5},
		{10, -1},
		{199, 9},
		{5, 8},
		{2, 3},
	}
	arry := []int{-1, 1, 1, 2, 2, 3, 3, 4, 5, 199, 200}
	for _, c := range cases {
		result := BinarySearchNest(arry, c.search)
		if result != c.want {
			t.Errorf("BinarySearchNest %d in %v, want %d, but return %d", c.search, arry, c.want, result)
		}
	}
}
