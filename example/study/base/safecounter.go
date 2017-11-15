package main

import (
	"sync"
	"time"
	"fmt"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (sc *SafeCounter) Inc(key string) {
	sc.mux.Lock()
	sc.v[key]++
	sc.mux.Unlock()
}

func (sc *SafeCounter) Value(key string) int {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	return sc.v[key]
}

func main() {
	sc := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go sc.Inc("somekey")
	}
	time.Sleep(3 * time.Second)
	fmt.Println(sc.Value("somekey"))
}
