package main

import "fmt"

func producer(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
	}
}

func consumer(n int, c chan int) int {
	result := 0
	for a := range c {
		result += a
	}
	fmt.Println("Result:", result)
	return result
}

func main() {
	channel := make(chan int)
	go producer(10, channel)
	go consumer(10, channel)
	// no output
	// prove that consumer does not receive any thing from producer
}
