package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// 1st <-c: assign c to x
	// 2st <-c: assign c to y
	// the magic is: 2 c are being scheduled, c is now a "fifo" queue
	x, y := <-c, <-c // receive from c
	// z := <-c // results in RTE (Runtime Error): fatal error: all goroutines are asleep - deadlock!

	fmt.Println(x, y, x+y)
}
