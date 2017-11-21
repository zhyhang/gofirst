package main

import (
	"time"
	"fmt"
)

func main() {
	// construct date using Date
	t:=time.Date(2017,11,17,11,12,13,888*1000000,time.Local)
	fmt.Println("Date time construct from Date():")
	printLeaddingTab(t)
}

func printLeaddingTab(t time.Time){
	fmt.Print("\t")
	fmt.Println(t)
}
