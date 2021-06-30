package main

import "fmt"

type Position struct {
	x float32
	y float32
}

func main() {
	personToPosition := make(map[string]Position)
	personToPosition["john sena"] = Position{1.2, 1.6}
	personToPosition["baker donought"] = Position{1.1, -1.6}
	personToPosition["lina vaccine"] = Position{10, -2}
	fmt.Println("People positions")
	fmt.Println(personToPosition)
}
