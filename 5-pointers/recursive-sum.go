package main

import "fmt"

func recursive_sum(arr []int) int {
	fmt.Printf("Address of arr is %p\n", &arr)
	if len(arr) == 1 {
		return arr[0]
	}
	return arr[0] + recursive_sum(arr[1:])
}
func main() {
	arr := []int{1, 2}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Address of element %d-th: %p\n", arr[i], &arr[i])
	}
	fmt.Println("Sum of", arr, "is", recursive_sum(arr))
}
