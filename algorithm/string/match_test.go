package string

import (
	"testing"
)

func TestSearchShortestSubString(t *testing.T) {
	cases := []struct {
		cs, letters, sub string
	}{
		{"", "Abc", ""},
		{"", "", ""},
		{"Hello world!", "", ""},
		{"Hello world!", "w", "w"},
		{"Hello world!", "H", "H"},
		{"Hello world!", "!", "!"},
		{"Hello world!", " ", " "},
		{"Hello world!", "Hello world! ", "Hello world!"},
		{"Hello world!", "lo", "lo"},
		{"Hello world!", "old", "orld"},
		{"Hello world!", "lHloo", "Hello"},
		{"Hello world!", "l!", "ld!"},
		{"Hello world!", "H!", "Hello world!"},
		{"Hello world!", "Abc", ""},
		{"Hello 世界!", "!世", "世界!"},
		{"Ab cAb  cAb  c", "cbA", "cAb"},
		{"AbcAbcAbc", "cbA", "Abc"},
		{"AbcAbcAbc", "def", ""},
	}

	for _, c := range cases {
		result := SearchShortestSubString(c.cs, c.letters)
		if result != c.sub {
			t.Errorf("SearchShortestSubString(%v,%v), want %v, but %v", c.cs, c.letters, c.sub, result)
		}
	}

}
