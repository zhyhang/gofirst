package sort

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

func QuickSortIter(a []int) {
	if a == nil || len(a) < 2 {
		return
	}
	stack := &istack{make([][2]int, len(a)), 0}
	stack.push([2]int{0, len(a) - 1})
	for !stack.isEmpty() {
		lowHigh := stack.pop()
		if lowHigh[0] >= lowHigh[1] {
			continue
		}
		loc := quickSortIterPartition(a, lowHigh[0], lowHigh[1])
		stack.push([2]int{loc + 1, lowHigh[1]})
		stack.push([2]int{lowHigh[0], loc - 1})
	}
}

func quickSortIterPartition(a []int, low, high int) int {
	pivot := a[high]
	i := low - 1
	for j := low; j < high; j++ {
		if a[j] <= pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[high] = pivot, a[i+1]
	return i + 1
}

type istack struct {
	values [][2]int
	top    int
}

func (s *istack) isEmpty() bool {
	return s.top == 0
}

func (s *istack) push(lowHigh [2]int) {
	s.values[s.top] = lowHigh
	s.top++
}

func (s *istack) pop() (lowHigh [2]int) {
	s.top--
	return s.values[s.top]
}
