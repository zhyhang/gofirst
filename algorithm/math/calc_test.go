package math

import (
	"testing"
)

func TestBigIntSubtract(t *testing.T) {
	cases := []struct {
		a, b, c []int
		sign    int
	}{
		{nil, []int{1, 2}, nil, 0},
		{[]int{3, 4}, nil, nil, 0},
		{nil, nil, nil, 0},
		{[]int{1, 1, 1}, []int{1, 1, 1}, []int{0, 0, 0}, 0},
		{[]int{0, 1, 1, 1}, []int{0, 0, 1, 1, 1}, []int{0, 0, 0, 0}, 0},
		{[]int{0, 0, 1, 1, 1}, []int{0, 1, 1, 1}, []int{0, 0, 0, 0, 0}, 0},
		{[]int{7, 8, 9}, []int{4, 5, 6}, []int{3, 3, 3}, 0},
		{[]int{4, 5, 6}, []int{7, 8, 9}, []int{3, 3, 3}, 1},
		{[]int{7, 7, 7}, []int{4, 8, 6}, []int{2, 9, 1}, 0},
		{[]int{6, 5, 4}, []int{4, 8, 6}, []int{1, 6, 8}, 0},
		{[]int{4, 8, 6}, []int{0, 7, 7, 7}, []int{0, 2, 9, 1}, 1},
		{[]int{0, 4, 8, 6}, []int{6, 5, 4}, []int{1, 6, 8}, 1},
		{[]int{6, 5, 4}, []int{0}, []int{6, 5, 4}, 0},
		{[]int{0}, []int{6, 5, 4}, []int{6, 5, 4}, 1},
		{[]int{6, 5, 4}, []int{}, []int{6, 5, 4}, 0},
		{[]int{}, []int{6, 5, 4}, []int{6, 5, 4}, 1},
		{[]int{6, 0, 4}, []int{3, 0, 8}, []int{2, 9, 6}, 0},
	}

	for _, c := range cases {
		result, sign, err := BigIntSubtract(c.a, c.b)
		if c.a == nil || c.b == nil {
			if err == nil {
				t.Errorf("BigIntSubtract(%v,%v), want error, but has no error", c.a, c.b)
			}
			continue
		}
		if sign != c.sign {
			t.Errorf("BigIntSubtract(%v,%v), want sign %v, but sign %v", c.a, c.b, c.sign, sign)
		}
		for i, r := range result {
			if r != c.c[i] {
				t.Errorf("BigIntSubtract(%v,%v), want %v, but %v", c.a, c.b, c.c, result)
			}
		}
	}
}
