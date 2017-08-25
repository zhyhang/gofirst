package search

import ()

//if found return index of e in a, or else return -1
func BinarySearch(a []int, e int) int {
	l, h := 0, len(a)-1
	for l <= h {
		m := (l + h) / 2
		if e == a[m] {
			return m
		} else if e < a[m] {
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

//if found return index of e in a, or else return -1
func BinarySearchNest(a []int, e int) int {
	return bsearch(a, 0, len(a)-1, e)
}

func bsearch(a []int, l, h, e int) int {
	if l > h {
		return -1
	}
	m := (l + h) / 2
	if e == a[m] {
		return m
	} else if e < a[m] {
		return bsearch(a, l, m-1, e)
	} else {
		return bsearch(a, m+1, h, e)
	}
}
