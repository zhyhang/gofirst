package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// float64 convert to duration
	//  type convert to time.Nanosecond
	seconds := math.Pi
	duration := time.Nanosecond * time.Duration(seconds*1e9)
	// duration convert to float64
	secondsNew := duration.Seconds()
	fmt.Printf("duration: %v, seconds value: %f\n", duration, secondsNew)
}
