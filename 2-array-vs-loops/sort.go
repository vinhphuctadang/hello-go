// print out array and sort descendingly
package main

import "fmt"

func main() {
	fmt.Printf("Array and for loop examples\n")

	var array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}

	fmt.Printf("\nArray after sorted DESC:\n")
	// sort the array desc
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] >= array[j] {
				continue
			}
			var tmp = array[i]
			array[i] = array[j]
			array[j] = tmp
		}
	}
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}
}
