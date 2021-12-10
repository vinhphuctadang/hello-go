package main

//#include "sum_array.h"
//#cgo LDFLAGS: sum_array.o
import "C"
import "fmt"

func main() {
	fmt.Println("Calling sum_array from  C code")
	a := []C.int{1, 2, 3, 4, 5}
	fmt.Println("Sum of a ==", C.sum_array(&a[0], C.int(len(a))))
}
