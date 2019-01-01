package list

import "testing"

func TestPrint(t *testing.T) {
	head := &slist{1, nil}
	for i, tail := 2, head; i <= 10; i, tail = i+1, tail.next {
		addNode(tail, i)
	}
	s := Print(head)
	t.Log(s)
}

func addNode(list *slist, value int) {
	list.next = &slist{value, nil}
}

func TestReverse(t *testing.T) {
	runReverse(t, Reverse)
}

func TestReverseNest(t *testing.T) {
	runReverse(t, ReverseNest)
}

func TestReverseInStack(t *testing.T) {
	runReverse(t, ReverseInStack)
}

func runReverse(t *testing.T, f func(*slist) *slist) {
	head := &slist{1, nil}
	for i, tail := 2, head; i <= 10; i, tail = i+1, tail.next {
		addNode(tail, i)
	}
	cases := []struct {
		sl        *slist
		wantPrint string
	}{
		{head, "head->10->9->8->7->6->5->4->3->2->1->nil"},
		{&slist{1, nil}, "head->1->nil"},
		{nil, "head->nil"},
	}

	for _, c := range cases {
		reversed := f(c.sl)
		revprint := Print(reversed)
		if revprint != c.wantPrint {
			t.Errorf("list: %s, want reversed: %s, but: %s", Print(c.sl), c.wantPrint, revprint)
		}
	}
}
