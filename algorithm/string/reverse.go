package string

import (
	dss "github.com/emirpasic/gods/stacks/arraystack"
)

// Reverse returns its argument string reversed rune-wise left to right
func Reverse(s string) string {
	r := []rune(s) // should convert to []rune for working well of unicode
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// reverse recursively
//
// e.g. the input string: abcdef
//   reverse(bcdef)+a
//   reverse(cdef)+b
//   ...
func ReverseNest(s string) string {
	r := []rune(s)
	if len(r) == 1 || len(r) == 0 {
		return s
	}
	sub := ReverseNest(string(r[1:]))
	return sub + string(r[0:1])
}

// reverse recursively (another manner)
func ReverseNest1(s string) string {
	r := []rune(s)
	switchHeadTail(r)
	return string(r)
}

func switchHeadTail(r []rune) {
	if len(r) <= 1 {
		return
	}
	bottom, top := 0, len(r)-1
	r[bottom], r[top] = r[top], r[bottom]
	switchHeadTail(r[bottom+1 : top])
}

// Reverse using stack
func ReverseInStack(s string) string {
	r := []rune(s)
	stack := dss.New()
	for i := 0; i < len(r); i++ {
		stack.Push(r[i])
	}
	rs := make([]rune, len(r))
	for i := 0; i < len(r); i++ {
		v, _ := stack.Pop()
		if v != nil {
			rs[i] = v.(rune)
		}
	}
	return string(rs)
}
