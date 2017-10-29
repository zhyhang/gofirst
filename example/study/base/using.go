package main

import (
	"fmt"
	"bytes"
)

func main() {
	// print
	fmt.Println("Hello world.\n")

	// point and value
	var good *string
	s := "Good morning!"
	good = &s

	if good != nil {
		fmt.Println(*good)
	}

	// map make and using
	map1 := make(map[string]bool)
	map1["s1"] = true
	map1["s2"] = false
	fmt.Println(map1)
	delete(map1,"s2")
	fmt.Println(map1)

	// new bytes.Buffer
	buf := new(bytes.Buffer)
	buf.WriteString("this is a testing of bytes.Buffer")
	fmt.Println(buf)
	var buf1 bytes.Buffer
	buf1.WriteString("this is another testing of bytes.Buffer")
	fmt.Println(&buf1)


}
