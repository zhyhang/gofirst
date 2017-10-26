package tree

import (
	"testing"
)

/*
func TestMirrorBinaryTree(t *testing.T) {
	cases :=[]struct{
		troot BiTree
		want string
	}{
		{BiTree{nil,nil,7788},""},
	}
	for _,c:=range cases {
		fmt.Println(TravelBinaryTree(&c.troot))
	}
}
*/

type treeCase struct {
	root    BiTree
	travels string
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

}
