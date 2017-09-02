package sort

import ()

// bubble sort a slice, asc
func Bubblesort(a []int) {
	for i, change := 0, true; i < len(a) && change; i++ {
		change = false
		for j := 1; j < len(a)-i; j++ {
			if a[j] < a[j-1] {
				a[j-1], a[j], change = a[j], a[j-1], true
			}
		}
	}
}

// quick sort a slice, asc
func Quicksort(a []int) {
	qsort(a, 0, len(a)-1)
}

func qsort(a []int, begin int, end int) {
	if begin < end {
		divIndex := qsortPartition(a, begin, end)
		qsort(a, begin, divIndex-1)
		qsort(a, divIndex+1, end)
	}
}

func qsortPartition(a []int, begin int, end int) int {
	pivot := a[begin]
	for begin < end {
		for ; a[end] >= pivot && begin < end; end-- {

		}
		if begin < end {
			a[begin] = a[end]
		}
		for ; a[begin] <= pivot && begin < end; begin++ {

		}
		if begin < end {
			a[end] = a[begin]
		}
		a[begin] = pivot
	}
	return begin
}

func Mergesort(a []int) {
	if len(a) == 0 {
		return
	}
	aux := make([]int, len(a))
	copy(aux, a)
	msort(aux, a, 0, len(a)-1)
}

func msort(src, dest []int, s, t int) {
	if s == t {
		dest[s] = src[s]
	} else {
		m := (s + t) / 2
		// Recursively sort
		msort(dest, src, s, m)
		msort(dest, src, m+1, t)
		msMerge(src, dest, s, m, t)
	}

}

// merge src[i,m],[m+1,n] to dest[i,max(m,n)]
func msMerge(src, dest []int, i, m, n int) {
	j, k := m+1, i
	for ; i <= m && j <= n; k++ {
		if src[i] <= src[j] {
			dest[k] = src[i]
			i++
		} else {
			dest[k] = src[j]
			j++
		}
	}
	for ; i <= m; i++ {
		dest[k] = src[i]
		k++
	}
	for ; j <= n; j++ {
		dest[k] = src[j]
		k++
	}
}
