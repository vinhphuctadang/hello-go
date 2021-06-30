package main

import "fmt"

func main() {
	/* = and := difference:
	 * = is assignment
	 * := is short-hand syntax of "declaration", ``a := 5`` means ``var a int = 5;``
	 */
	arr := []int{1, 2}
	fmt.Println("Array:", arr)
	fmt.Printf("Array address: %p\n", &arr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Address of element %d-th: %p\n", arr[i], &arr[i])
	}

	arr = append(arr[:], 3)
	fmt.Println("Array after appended:", arr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Address of element %d-th: %p\n", arr[i], &arr[i])
	}
}
