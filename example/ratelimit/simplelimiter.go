package main

import (
	"fmt"
	"math/rand"
)

import "time"

type Request interface{}

var bTime = time.Now()

func handle(r Request) {
	if rand.Float32() > 0.318 {
		elapseSeconds := time.Now().Sub(bTime).Seconds()
		if elapseSeconds > 0 {
			c := r.(int)
			fmt.Printf("current count: %d, qps: %f\n", c, float64(c)/elapseSeconds)
		}
	}
}

const RateLimitPeriod = time.Minute
const RateLimit = 600 // most 600 requests in one minute

func handleRequests(requests chan Request) {
	quotas := make(chan time.Time, RateLimit)

	go func() {
		tick := time.NewTicker(RateLimitPeriod / RateLimit)
		defer tick.Stop()
		for t := range tick.C {
			select {
			case quotas <- t:
			default:
			}
		}
	}()

	for r := range requests {
		<-quotas
		go handle(r)
	}
}

func main() {
	requests := make(chan Request)
	go handleRequests(requests)

	for i := 0; ; i++ {
		requests <- i
	}
}
