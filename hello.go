package main

import (
	"fmt"
	"github.com/zhyhang/gofirst/stringutil"
)

func main() {
	fmt.Println("Hello world.\n")
	fmt.Println(stringutil.Reverse("Hello world."))
	fmt.Println(stringutil.Reverse("Good morning."))
}
