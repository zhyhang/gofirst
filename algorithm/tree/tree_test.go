package tree

import (
	"testing"
)

func TestMirrorBinaryTree(t *testing.T) {
	// nil tree
	root := NewBinaryTreeNode("", "", "")
	MirrorBinaryTree(root)
	rootStr := TravelBinaryTree(root)
	if "" != rootStr {
		t.Errorf("MirrorBinaryTree, want %s, but %s", "nil", rootStr)
	}
	// most 3 nodes
	cases := []struct {
		o, l, r, want string
	}{
		{"1", "2", "3", "1,3,2"},
		{"1", "", "3", "1,3"},
		{"1", "2", "", "1,2"},
	}
	for _, c := range cases {
		root = NewBinaryTreeNode(c.o, c.l, c.r)
		MirrorBinaryTree(root)
		rootStr = TravelBinaryTree(root)
		if c.want != rootStr {
			t.Errorf("MirrorBinaryTree, want %s, but %s", c.want, rootStr)
		}
	}
	// only have many left nodes
	root = NewBinaryTreeNode("1", "2", "")
	node := NewBinaryTreeNode("10", "20", "")
	root.left.left = node
	node.left.left = NewBinaryTreeNode("100", "200", "")
	MirrorBinaryTree(root)
	rootStr = TravelBinaryTree(root)
	if "1,2,10,20,100,200" != rootStr {
		t.Errorf("MirrorBinaryTree, want %s, but %s", "1,2,10,20,100,200", rootStr)
	}
	// only have many right nodes
	root = NewBinaryTreeNode("1", "", "3")
	node = NewBinaryTreeNode("10", "", "30")
	root.right.right = node
	node.right.right = NewBinaryTreeNode("100", "", "300")
	MirrorBinaryTree(root)
	rootStr = TravelBinaryTree(root)
	if "1,3,10,30,100,300" != rootStr {
		t.Errorf("MirrorBinaryTree, want %s, but %s", "1,3,10,30,100,300", rootStr)
	}
	// height=3, full binary tree (lack of right in bottom level)
	root = NewBinaryTreeNode("1", "2", "3")
	node = NewBinaryTreeNode("10", "20", "30")
	root.left = node
	MirrorBinaryTree(root)
	rootStr = TravelBinaryTree(root)
	if "1,3,10,30,20" != rootStr {
		t.Errorf("MirrorBinaryTree, want %s, but %s", "1,3,10,30,20", rootStr)
	}
	// height=3, full binary tree (lack of left in bottom level)
	root = NewBinaryTreeNode("1", "2", "3")
	node = NewBinaryTreeNode("10", "20", "30")
	root.right = node
	MirrorBinaryTree(root)
	rootStr = TravelBinaryTree(root)
	if "1,10,30,20,2" != rootStr {
		t.Errorf("MirrorBinaryTree, want %s, but %s", "1,10,30,20,2", rootStr)
	}
	// height=3, complete tree
	root = NewBinaryTreeNode("1", "2", "3")
	node = NewBinaryTreeNode("2", "21", "22")
	root.left = node
	node = NewBinaryTreeNode("3", "31", "32")
	root.right = node
	MirrorBinaryTree(root)
	rootStr = TravelBinaryTree(root)
	if "1,3,32,31,2,22,21" != rootStr {
		t.Errorf("MirrorBinaryTree, want %s, but %s", "1,3,32,31,2,22,21", rootStr)
	}
	// any tree
	root = NewBinaryTreeNode("1", "2", "5")
	node = NewBinaryTreeNode("2", "", "3")
	root.left = node
	root.right = NewBinaryTreeNode("5", "6", "7")
	node.right = NewBinaryTreeNode("3", "4", "")
	MirrorBinaryTree(root)
	rootStr = TravelBinaryTree(root)
	if "1,5,7,6,2,3,4" != rootStr {
		t.Errorf("MirrorBinaryTree, want %s, but %s", "1,5,7,6,2,3,4", rootStr)
	}
}

func TestTravelBinaryTree(t *testing.T) {
	// nil tree
	root := NewBinaryTreeNode("", "", "")
	rootStr := TravelBinaryTree(root)
	if "" != rootStr {
		t.Errorf("TravelBinaryTree, want %s, but %s", "nil", rootStr)
	}
	// most 3 nodes
	cases := []struct {
		o, l, r, want string
	}{
		{"1", "2", "3", "1,2,3"},
		{"1", "", "3", "1,3"},
		{"1", "2", "", "1,2"},
	}
	for _, c := range cases {
		root = NewBinaryTreeNode(c.o, c.l, c.r)
		rootStr = TravelBinaryTree(root)
		if c.want != rootStr {
			t.Errorf("TravelBinaryTree, want %s, but %s", c.want, rootStr)
		}
	}
	// only have many left nodes
	root = NewBinaryTreeNode("1", "2", "")
	node := NewBinaryTreeNode("10", "20", "")
	root.left.left = node
	node.left.left = NewBinaryTreeNode("100", "200", "")
	rootStr = TravelBinaryTree(root)
	if "1,2,10,20,100,200" != rootStr {
		t.Errorf("TravelBinaryTree, want %s, but %s", "1,2,10,20,100,200", rootStr)
	}
	// only have many  right nodes
	root = NewBinaryTreeNode("1", "", "3")
	node = NewBinaryTreeNode("10", "", "30")
	root.right.right = node
	node.right.right = NewBinaryTreeNode("100", "", "300")
	rootStr = TravelBinaryTree(root)
	if "1,3,10,30,100,300" != rootStr {
		t.Errorf("TravelBinaryTree, want %s, but %s", "1,3,10,30,100,300", rootStr)
	}
	// height=3, full binary tree (lack of right in bottom level)
	root = NewBinaryTreeNode("1", "2", "3")
	node = NewBinaryTreeNode("10", "20", "30")
	root.left = node
	rootStr = TravelBinaryTree(root)
	if "1,10,20,30,3" != rootStr {
		t.Errorf("TravelBinaryTree, want %s, but %s", "1,10,20,30,3", rootStr)
	}
	// height=3, full binary tree (lack of left in bottom level)
	root = NewBinaryTreeNode("1", "2", "3")
	node = NewBinaryTreeNode("10", "20", "30")
	root.right = node
	rootStr = TravelBinaryTree(root)
	if "1,2,10,20,30" != rootStr {
		t.Errorf("TravelBinaryTree, want %s, but %s", "1,2,10,20,30", rootStr)
	}
	// height=3, complete tree
	root = NewBinaryTreeNode("1", "2", "3")
	node = NewBinaryTreeNode("2", "21", "22")
	root.left = node
	node = NewBinaryTreeNode("3", "31", "32")
	root.right = node
	rootStr = TravelBinaryTree(root)
	if "1,2,21,22,3,31,32" != rootStr {
		t.Errorf("TravelBinaryTree, want %s, but %s", "1,2,21,22,3,31,32", rootStr)
	}
	// any tree
	root = NewBinaryTreeNode("1", "2", "5")
	node = NewBinaryTreeNode("2", "", "3")
	root.left = node
	root.right = NewBinaryTreeNode("5", "6", "7")
	node.right = NewBinaryTreeNode("3", "4", "")
	rootStr = TravelBinaryTree(root)
	if "1,2,3,4,5,6,7" != rootStr {
		t.Errorf("TravelBinaryTree, want %s, but %s", "1,2,3,4,5,6,7", rootStr)
	}
}

func TestNewBinaryTreeNode(t *testing.T) {
	// normal input
	cases := []struct {
		root, left, right    string
		iroot, ileft, iright int
	}{
		{"1", "2", "3", 1, 2, 3},
		{"4", "5", "6", 4, 5, 6},
		{"-2017", "1978", "12", -2017, 1978, 12},
	}
	for _, c := range cases {
		node := NewBinaryTreeNode(c.root, c.left, c.right)
		if node.value != c.iroot || node.left.value != c.ileft || node.right.value != c.iright {
			t.Errorf("NewBinaryTreeNode, want (%d,%d,%d), but (%d,%d,%d)", c.iroot, c.ileft, c.iright, node.value, node.left.value, node.right.value)
		}
	}
	// error input
	node := NewBinaryTreeNode("a", "1", "2")
	if node != nil {
		t.Errorf("NewBinaryTreeNode(a,1,2), want nil, but %v", node)
	}
	node = NewBinaryTreeNode("-1024", "", "2")
	if node.left != nil || node.right == nil || node.value != -1024 || node.right.value != 2 {
		t.Errorf("NewBinaryTreeNode(-1024,nil,2), want (-1024,nil,2), but %v", node)
	}
	node = NewBinaryTreeNode("-1024", "1", "")
	if node.right != nil || node.left == nil || node.value != -1024 || node.left.value != 1 {
		t.Errorf("NewBinaryTreeNode(-1024,nil,2), want (-1024,1,nil), but %v", node)
	}
	node = NewBinaryTreeNode("-1024", "error1", "")
	if node.right != nil || node.left != nil || node.value != -1024 {
		t.Errorf("NewBinaryTreeNode(-1024,error1,nil), want (-1024,nil,nil), but %v", node)
	}
}
