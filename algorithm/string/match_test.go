package string

import (
	"reflect"
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

func TestIndexOfAll(t *testing.T) {
	cases := []struct {
		txt, sub string
		indexes  []int
	}{
		{"", "", []int{0}},
		{"a", "", nil},
		{"abcdefga", "a", []int{0, 7}},
		{"abcdefga", "c", []int{2}},
		{"aaabcaadaa", "aa", []int{0, 1, 5, 8}},
		{"abcdefga", "ac", nil},
	}

	for _, c := range cases {
		result := IndexOfAll(c.txt, c.sub)
		if !reflect.DeepEqual(result, c.indexes) {
			t.Errorf("IndexOfAll(%v,%v), want %v, but %v", c.txt, c.sub, c.indexes, result)
		}
	}

	// test for simple

	for _, c := range cases {
		result := IndexOfAllSimple(c.txt, c.sub)
		if !reflect.DeepEqual(result, c.indexes) {
			t.Errorf("IndexOfAllSimple(%v,%v), want %v, but %v", c.txt, c.sub, c.indexes, result)
		}
	}

}
