package string

import "testing"

func TestReverse(t *testing.T) {
	runReverse(t, Reverse)
}

func TestReverseNest(t *testing.T) {
	runReverse(t, ReverseNest)
}

func TestReverseInStack(t *testing.T) {
	runReverse(t, ReverseInStack)
}

func runReverse(t *testing.T, workingMethod func(string) string) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := workingMethod(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
