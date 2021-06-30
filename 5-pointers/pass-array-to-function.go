package main

import "fmt"

func subtract_each(arr []int, by int) {
	fmt.Printf("Array address inside function: %p\n", &arr)
	for i := 0; i < len(arr); i++ {
		arr[i] -= by
	}
}

func main() {
	arr := []int{1, 3, -4, 6, 9}
	fmt.Printf("Array address: %p\n", &arr)
	fmt.Println("Array:", arr)
	subtract_each(arr[:], 1) // a[i] -= 1s
	fmt.Println("Array after subtraction:", arr)
}
