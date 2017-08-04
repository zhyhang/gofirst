package sort

import ()

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

func Quicksort(a []int) {

}

func qsort(a []int, begin int, end int) {

}

func qsortpartition(a []int, begin int, end int) {

}
