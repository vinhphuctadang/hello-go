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
	fmt.Printf("Done computing sum of %d integers, result = %d\n", n, result)
}

func main() {
	go sum(100000000)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Finished")
}
