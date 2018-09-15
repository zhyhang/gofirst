package json_test

import (
	"github.com/zhyhang/gofirst/example/study/base/json"
	"testing"
)

func TestOfInt(t *testing.T) {
	json.OfInt(1024)
	json.OfInt(-1024)
}

func TestOfSlice(t *testing.T) {
	json.OfSlice([]string{"hello world!"})
}

func TestOfMap(t *testing.T) {
	json.OfMap(map[string]int{"s1": 10, "s2": 20})
}

func TestOfStruct(t *testing.T) {
	type fort struct {
		V1 int    `json:"age"`
		V2 string `json:"name"`
	}
	st1 := &fort{
		18, "yu",
	}
	json.OfStruct(st1)
}
