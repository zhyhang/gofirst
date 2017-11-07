package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, av := range a {
		sum += av
	}
	c <- sum
}

func main() {
	c := make(chan int)
	a := []int{1,3, 5, 7, 8, 10, 12}
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Printf("x=%d,y=%d,sum=%d\n", x, y, x+y)
}
