package sort

import (
	"fmt"
	"testing"
)

func TestBubblesort(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{3, 5, -1, 0, 6, 4}, []int{-1, 0, 3, 4, 5, 6}},
		{[]int{-3, -2, -1, 0, 1, 2}, []int{-3, -2, -1, 0, 1, 2}},
		{make([]int, 0, 0), make([]int, 0, 0)},
	}
	for _, c := range cases {
		Bubblesort(c.in)
		for i := 0; i < len(c.in); i++ {
			if c.in[i] != c.want[i] {
				t.Errorf("Bubblesort(%v), want %v", c.in, c.want)
			}
		}
	}
}

func TestQuicksort(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{3, 5, -1, 0, 6, 4}, []int{-1, 0, 3, 4, 5, 6}},
		{[]int{-3, -2, -1, 0, 1, 2}, []int{-3, -2, -1, 0, 1, 2}},
		{[]int{3, 2, 1, 4, -1, 200, 199, 1, 2, 3, 5}, []int{-1, 1, 1, 2, 2, 3, 3, 4, 5, 199, 200}},
		{make([]int, 0, 0), make([]int, 0, 0)},
	}
	for _, c := range cases {
		fmt.Println(c.in)
		Quicksort(c.in)
		fmt.Println(c.in)
		for i, e := range c.in {
			if e != c.want[i] {
				t.Errorf("Quicksort(%v), want %v", c.in, c.want)
			}
		}
	}
}

func TestMergesort(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{3, 5, -1, 0, 6, 4}, []int{-1, 0, 3, 4, 5, 6}},
		{[]int{-3, -2, -1, 0, 1, 2}, []int{-3, -2, -1, 0, 1, 2}},
		{[]int{3, 2, 1, 4, -1, 200, 199, 1, 2, 3, 5}, []int{-1, 1, 1, 2, 2, 3, 3, 4, 5, 199, 200}},
		{make([]int, 0, 0), make([]int, 0, 0)},
	}
	for _, c := range cases {
		fmt.Println(c.in)
		Mergesort(c.in)
		fmt.Println(c.in)
		for i, e := range c.in {
			if e != c.want[i] {
				t.Errorf("Mergesort(%v), want %v", c.in, c.want)
			}
		}
	}
}

func TestQuickSortIter(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{3, 5, -1, 0, 6, 4}, []int{-1, 0, 3, 4, 5, 6}},
		{[]int{-3, -2, -1, 0, 1, 2}, []int{-3, -2, -1, 0, 1, 2}},
		{[]int{3, 2, 1, 4, -1, 200, 199, 1, 2, 3, 5}, []int{-1, 1, 1, 2, 2, 3, 3, 4, 5, 199, 200}},
		{make([]int, 0, 0), make([]int, 0, 0)},
	}
	for _, c := range cases {
		fmt.Println(c.in)
		QuickSortIter(c.in)
		fmt.Println(c.in)
		for i, e := range c.in {
			if e != c.want[i] {
				t.Errorf("QuickSortIter(%v), want %v", c.in, c.want)
			}
		}
	}
}
