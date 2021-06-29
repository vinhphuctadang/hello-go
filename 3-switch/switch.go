package main

import "fmt"

// rune is alias of int32

// func function_name(name type, name type, ...) return_type/(return_typ1, return_typ2)
func number_of_days(month rune, year rune) rune {
	switch month {
	case 2:
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			return 29
		}
		return 28
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
}

func main() {
	month := 4
	year := 1900
	fmt.Printf("%d/%d has %d days", month, year, number_of_days(rune(month), rune(year)))
}
