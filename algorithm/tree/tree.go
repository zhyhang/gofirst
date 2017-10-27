package tree

import (
	"bytes"
	"fmt"
	"strconv"
)

type BiTree struct {
	left, right *BiTree
	value       int
}

func MirrorBinaryTree(root *BiTree) {
	if root == nil {
		return
	}
	root.left, root.right = root.right, root.left
	MirrorBinaryTree(root.left)
	MirrorBinaryTree(root.right)
}

func TravelBinaryTree(root *BiTree) string {
	if root == nil {
		return ""
	}
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%d,", root.value)
	travelBinaryTree(buf, root.left)
	travelBinaryTree(buf, root.right)
	bufStr := buf.String()
	return bufStr[0 : len(bufStr)-1]

}

func travelBinaryTree(buf *bytes.Buffer, root *BiTree) {
	if root != nil {
		fmt.Fprintf(buf, "%d,", root.value)
		travelBinaryTree(buf, root.left)
		travelBinaryTree(buf, root.right)
	}
}

// make a node of binary tree with left and right nodes
// root should not nil, left and right can nil
func NewBinaryTreeNode(root string, left string, right string) *BiTree {
	rootVal, rootErr := strconv.Atoi(root)
	leftVal, leftErr := strconv.Atoi(left)
	rightVal, rightErr := strconv.Atoi(right)
	if rootErr != nil {
		return nil
	}
	rootNode := BiTree{nil, nil, rootVal}
	if leftErr == nil {
		rootNode.left = &BiTree{nil, nil, leftVal}
	}
	if rightErr == nil {
		rootNode.right = &BiTree{nil, nil, rightVal}
	}
	return &rootNode
}
