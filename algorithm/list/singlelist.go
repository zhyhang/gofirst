package list

import (
	"bytes"
	"fmt"
	dss "github.com/emirpasic/gods/stacks/arraystack"
)

type slist struct {
	value int
	next  *slist
}

// the list: head->1->2
//
// before 1 iteration:
//  current=1
//  pre=nil
//  next=nil
//  1.next=2
// after 1 iteration:
//  current=2
//  pre=1
//  next=2
//  1.next=nil
// after 2 iteration:
//  current=nil
//  pre=2
//  next=nil
//  2.next=1
func Reverse(list *slist) *slist {
	var prev, next *slist
	current := list
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	return prev
}

// recursively reverse
func ReverseNest(list *slist) *slist {
	if list == nil {
		return list
	}
	return reverseNestInner(list, nil)
}

func reverseNestInner(head, prev *slist) *slist {
	next := head.next
	head.next = prev
	if next == nil {
		return head
	}
	return reverseNestInner(next, head)
}

func ReverseInStack(list *slist) *slist {
	if list == nil || list.next == nil {
		return list
	}
	stack := dss.New()
	// push to stack
	for node := list; node != nil; node = node.next {
		stack.Push(node)
	}
	head, _ := stack.Pop()
	prev := head.(*slist)
	// pop from stack
	for prev != nil {
		n, _ := stack.Pop()
		node, ok := n.(*slist)
		if ok {
			prev.next = node
			prev = node
		} else {
			prev.next = nil
			prev = nil
		}
	}

	return head.(*slist)

}

func Print(list *slist) string {
	buf := new(bytes.Buffer)
	inList := list
	fmt.Fprint(buf, "head->")
	for ; inList != nil; inList = inList.next {
		fmt.Fprint(buf, inList.value)
		fmt.Fprint(buf, "->")
	}
	fmt.Fprint(buf, "nil")
	return buf.String()
}
