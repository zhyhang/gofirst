package sort

import (
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
