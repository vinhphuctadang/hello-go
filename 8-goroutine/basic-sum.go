package main

import (
	"fmt"
	"time"
)

// assume that we compute sum of first N integers using For-loop

func sum(n int) {
	result := 0
	for i := 0; i < n; i++ {
		result += i
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Done computing sum of %d integers, result = %d\n", n, result)
}

func main() {
	go sum(100000000)
	sum(1000)
	fmt.Printf("Finished")
}
