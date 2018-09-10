package main

import (
	"fmt"
	"math"
)

// concurrent computation of π
// This demonstrates Go's ability to handle
// large numbers of concurrent processes.
func main() {
	iterates := 5000
	ch := make(chan float64)
	for i := 0; i <= iterates; i++ {
		go computerTerm(ch, float64(i))
	}
	pi := 0.0
	for i := 0; i <= iterates; i++ {
		pi += <-ch
	}
	fmt.Printf("π is: %f \n", pi)
}

func computerTerm(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / float64(2*k+1)
}
