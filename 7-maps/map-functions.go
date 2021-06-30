package main

import "fmt"

type PersonRecord struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type RecordApplicableFunction func(record PersonRecord)

func main() {
	funcMap := map[string]RecordApplicableFunction{
		"print": func(record PersonRecord) {
			fmt.Println("Person:", record)
		},
	}
	funcMap["print"](PersonRecord{1, "vinhphuctadang", 22})
}
