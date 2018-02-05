package main

import (
	"fmt"
	"github.com/xluohome/phonedata"
)

func main() {
	pr, err := phonedata.Find("13513258895")
	if err != nil {
		panic(err)
	}
	fmt.Print(pr)
}
