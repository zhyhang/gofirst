package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		search, want int
	}{
		{-1, 0},
		{200, 10},
		{4, 7},
		{3, 5},
		{10, -1},
		{199, 9},
		{5, 8},
		{2, 3},
	}
	arry := []int{-1, 1, 1, 2, 2, 3, 3, 4, 5, 199, 200}
	for _, c := range cases {
		result := BinarySearch(arry, c.search)
		if result != c.want {
			t.Errorf("BinarySearch %d in %v, want %d, but return %d", c.search, arry, c.want, result)
		}
	}
}

func TestBinarySearchNest(t *testing.T) {
	cases := []struct {
		search, want int
	}{
		{-1, 0},
		{200, 10},
		{4, 7},
		{3, 5},
		{10, -1},
		{199, 9},
		{5, 8},
		{2, 3},
	}
	arry := []int{-1, 1, 1, 2, 2, 3, 3, 4, 5, 199, 200}
	for _, c := range cases {
		result := BinarySearchNest(arry, c.search)
		if result != c.want {
			t.Errorf("BinarySearchNest %d in %v, want %d, but return %d", c.search, arry, c.want, result)
		}
	}
}

func TestBinarySearchNearest(t *testing.T) {

	v := 0

	//# root node is nil
	result := BinarySearchNearest(nil, v)
	if result != nil {
		t.Errorf("BinarySearchNearest %d in nil, want nil, but return %v", v, result)
	}

	//# only one root node
	root := &BiTreeNode{
		nil, nil, 10,
	}

	// root.value==v
	v = 10
	result = BinarySearchNearest(root, v)
	if result != root {
		t.Errorf("BinarySearchNearest %d in %v, want %v, but return %v", v, root, root, result)
	}

	// root.value<v
	v = 11
	result = BinarySearchNearest(root, v)
	if result != root {
		t.Errorf("BinarySearchNearest %d in %v, want %v, but return %v", v, root, root, result)
	}

	// root.value>v
	v = 4
	result = BinarySearchNearest(root, v)
	if result != root {
		t.Errorf("BinarySearchNearest %d in %v, want %v, but return %v", v, root, root, result)
	}

	//# root has left child and right child
	//
	//				10
	//		  8				12
	//	7			9					14
	//							13
	//
	left := &BiTreeNode{
		nil, nil, 7,
	}

	right := &BiTreeNode{
		nil, nil, 9,
	}

	rootLeft := &BiTreeNode{
		left, right, 8,
	}

	left = &BiTreeNode{
		nil, nil, 13,
	}
	right = &BiTreeNode{
		left, nil, 14,
	}

	rootRight := &BiTreeNode{
		nil, right, 12,
	}

	root = &BiTreeNode{
		rootLeft, rootRight, 10,
	}

	// equal
	for v = 7; v <= 10; v++ {
		result = BinarySearchNearest(root, v)
		if v != result.value {
			t.Errorf("BinarySearchNearest %d in %v, want %d, but return %d", v, root, v, result.value)
		}
	}

	// near
	v = 11
	result = BinarySearchNearest(root, v)
	if result != rootRight {
		t.Errorf("BinarySearchNearest %d in %v, want %v, but return %v", v, root, rootRight, result)
	}

	// near
	v = 4
	result = BinarySearchNearest(root, v)
	if 7 != result.value {
		t.Errorf("BinarySearchNearest %d in %v, want %d, but return %d", v, root, 7, result.value)
	}

	// near
	v = 15
	result = BinarySearchNearest(root, v)
	if 14 != result.value {
		t.Errorf("BinarySearchNearest %d in %v, want %d, but return %d", v, root, 15, result.value)
	}

	// equal
	for v = 12; v <= 14; v++ {
		result = BinarySearchNearest(root, v)
		if v != result.value {
			t.Errorf("BinarySearchNearest %d in %v, want %d, but return %d", v, root, v, result.value)
		}
	}
}
