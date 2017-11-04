package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go say("world")
	say("hello")
	wg.Wait()// wait done the goroutine
}
