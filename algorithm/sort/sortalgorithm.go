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
