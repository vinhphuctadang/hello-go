package main

import "fmt"

func main() {
	a := int32(10)
	b := int32(20)
	arrA := [100]int32{}
	// arrB := [100]int32{}

	pnt_a := &a
	var pnt_arr *[]int32

	fmt.Printf("Address of a: %p\n", pnt_a)
	fmt.Printf("Address of b: %p\n", &b)

	fmt.Printf("Value of a : %d\n", a)
	fmt.Printf("Value of b : %d\n", b)

	fmt.Printf("Address of arrA: %p\n", &arrA)
	fmt.Println(arrA)

	pnt_a = &b
	*pnt_a = a
	fmt.Printf("Address referred by pnt_a: %p\n", pnt_a)
	fmt.Printf("Value of b: %d\n", b)

	// pass array to function, use slice

	// Go looks not support this way, may be its array type has "length" stored :D :D
	// pnt_arr = &arrA
	// pnt_arr[0] = 100
	// fmt.Printf("Address referred by pnt_arr: %p\n", pnt_arr)
	// fmt.Println(pnt_arr)

	// Go does NOT support pointer arithmetic which is present in other languages like C and C++. https://golangbot.com/pointers/
}
