package main

import (
	"fmt"
	"bytes"
	"strconv"
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
	fmt.Println()

	// map make and using
	map1 := make(map[string]bool)
	map1["s1"] = true
	map1["s2"] = false
	fmt.Println(map1)
	delete(map1, "s2")
	fmt.Println(map1)
	fmt.Println()

	// new bytes.Buffer
	buf := new(bytes.Buffer)
	buf.WriteString("this is a testing of bytes.Buffer")
	fmt.Println(buf)
	var buf1 bytes.Buffer
	buf1.WriteString("this is another testing of bytes.Buffer")
	fmt.Println(&buf1)
	fmt.Println()

	// channel
	ch1 := make(chan string)
	for i := 0; i < 10; i++ {
		go func(c int) {
			ch1 <- strconv.Itoa(c)
		}(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch1)
	}
	fmt.Println()

}
