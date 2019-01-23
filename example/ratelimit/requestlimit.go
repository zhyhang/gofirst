package main

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"sync/atomic"
	"time"
)

func main() {
	limiter := rate.NewLimiter(rate.Limit(1000), 10000)
	var total int32
	beginTime := time.Now()
	waitChannel := make(chan int, 1)
	for i := 0; i < 24; i++ {
		go task(&total, &beginTime, limiter)
	}
	<-waitChannel

}

func task(total *int32, beginTime *time.Time, limiter *rate.Limiter) {
	bTime := *beginTime
	lastTotal := *total
	for {
		limiter.Wait(context.TODO())
		atomic.AddInt32(total, 1)
		timeNow := time.Now()
		subTime := timeNow.Sub(bTime)
		takenSeconds := subTime.Seconds()
		if takenSeconds >= 1.0 {
			bTime = timeNow
			currentTotal := *total
			totalDiff := currentTotal - lastTotal
			lastTotal = currentTotal
			log.Printf("current total: %d, qps: %f", currentTotal, float64(totalDiff)/takenSeconds)
		}
	}
}
