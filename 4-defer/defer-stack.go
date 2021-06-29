package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		defer fmt.Print("This is ", i, " run when main finished\n")
	}
	fmt.Println("Going to do some stuffs")
}
