package main

import (
	"fmt"
)

func main() {
	var number, lastDigit, currentDigit int
	var isAscending, isDescending bool

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&number)

	isAscending = true
	isDescending = true

	// mengambil digit
	lastDigit = number % 10
	number /= 10

	for number > 0 {
		currentDigit = number % 10
		if currentDigit > lastDigit {
			isAscending = false
		} else if currentDigit < lastDigit {
			isDescending = false
		}
		lastDigit = currentDigit
		number /= 10
	}

	if isAscending {
		fmt.Println("ascending")
	} else if isDescending {
		fmt.Println("descending")
	} else {
		fmt.Println("tidak terurut")
	}
}
