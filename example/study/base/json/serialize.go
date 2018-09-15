package json

import (
	"encoding/json"
	"log"
)

func OfInt(i int) {
	jsonBytes, _ := json.Marshal(i)
	log.Printf("json of int %d is %v\n", i, string(jsonBytes))
}

func OfSlice(ss []string) {
	jsonBytes, _ := json.Marshal(ss)
	log.Printf("json of slice %v is\n %s\n", ss, string(jsonBytes))

}

func OfMap(m map[string]int) {
	jsonBytes, _ := json.Marshal(m)
	log.Printf("json of %v is\n %s\n", m, string(jsonBytes))
}

func OfStruct(st interface{}) {
	jsonBytes, _ := json.Marshal(st)
	log.Printf("json of %v is\n %s\n", st, string(jsonBytes))
}
