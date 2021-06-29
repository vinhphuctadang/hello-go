package main

import "fmt"

func main() {
	defer fmt.Print("This is run when main finished")
	fmt.Println("Going to do some stuffs")
}
