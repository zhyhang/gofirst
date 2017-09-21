package search

import (
	"math"
)

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

type BiTreeNode struct {
	left, right *BiTreeNode
	value       int
}

// find the nearest node value in binary search tree
func BinarySearchNearest(root *BiTreeNode, v int) *BiTreeNode {
	if root == nil {
		return nil
	}
	if root.value == v {
		return root
	}
	if v < root.value {
		node := BinarySearchNearest(root.left, v)
		if node == nil || math.Abs((float64)(v-node.value)) > math.Abs((float64)(v-root.value)) {
			return root
		} else {
			return node
		}
	}
	if v > root.value {
		node := BinarySearchNearest(root.right, v)
		if node == nil || math.Abs((float64)(v-node.value)) > math.Abs((float64)(v-root.value)) {
			return root
		} else {
			return node
		}
	}
	return nil
}
