package main

import (
	"fmt"
)

type PersonRecord struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (record *PersonRecord) requestPrint() {
	fmt.Println("Person:", record)
}

func main() {
	people := []PersonRecord{
		{1, "John Adam", 21},
		{2, "Wick Laugh", 29},
		{3, "Canoe Boat", 48},
	}

	// declare with type is ok but no need (PersonRecord)
	// people := []PersonRecord{
	// 	PersonRecord{1, "John Adam", 21},
	// 	PersonRecord{2, "Wick Laugh", 29},
	// 	PersonRecord{3, "Canoe Boat", 48},
	// }

	fmt.Println("Collected person infos:", people)

	people[0].requestPrint()

}
