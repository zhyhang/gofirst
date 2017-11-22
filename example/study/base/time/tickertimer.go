package main

import (
	"time"
	"fmt"
)

func main() {

	// sleep
	fmt.Printf("current millis: %d\n", time.Now().UnixNano()/1000000)
	time.Sleep(618 * time.Millisecond)
	fmt.Printf("current millis: %d, after sleep\n", time.Now().UnixNano()/1000000)

	// ticker (can repeatly send envents)
	i := 0
	for tkt := range time.Tick(time.Millisecond * 100) {
		if i < 10 {
			fmt.Printf("tick time: %s\n", tkt.Format(time.StampMilli))
		} else {
			break
		}
		i++
	}

	// new ticker
	tkr := time.NewTicker(time.Millisecond * 100)
	i = 0
	for tkt := range tkr.C {
		if i < 10 {
			fmt.Printf("new ticker tick time: %s\n", tkt.Format(time.StampMilli))
		} else {
			tkr.Stop()
			break
		}
		i++
	}

	// timer (only send event once)
	timer := time.NewTimer(time.Millisecond * 100)
	tmt := <-timer.C
	fmt.Printf("timer time: %s", tmt.Format(time.StampMilli))
	timer.Stop()

}
