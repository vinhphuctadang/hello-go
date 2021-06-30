package main

import (
	"encoding/json"
	"fmt"
)

type PersonRecord struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	people := []PersonRecord{
		{1, "John Adam", 21},
		{2, "Wick Laugh", 29},
		{3, "Canoe Boat", 48},
	}

	// additional type clarify is ok but no need (PersonRecord)
	// people := []PersonRecord{
	// 	PersonRecord{1, "John Adam", 21},
	// 	PersonRecord{2, "Wick Laugh", 29},
	// 	PersonRecord{3, "Canoe Boat", 48},
	// }

	fmt.Println("Collected person infos:", people)

	// make pointer to first element
	var john *PersonRecord = &people[0]
	john.Name = "John Beckam"

	fmt.Println("Collected person infos after changes:", people)

	jsonData, _ := json.Marshal(people)
	fmt.Println("Collected person infos in json:", string(jsonData))
}
