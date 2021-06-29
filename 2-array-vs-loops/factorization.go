// factorization for a natural number n

package main

import "fmt"

func main() {
	n := 16042610
	current_divisor := 2

	// Says, it's time to drop "while" loop, there is only "For" in Go
	for n > 1 {
		power := 0
		// find next divisor of current n
		for n%current_divisor != 0 {
			current_divisor++
		}
		// compute how many times can ''current_divisor'' can be multiplied
		for n%current_divisor == 0 {
			power++
			n /= current_divisor
		}
		fmt.Printf("%d^%d ", current_divisor, power)
	}
}
