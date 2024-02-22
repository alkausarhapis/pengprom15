package main

import "fmt"

func main() {
	var m, n int

	fmt.Scan(&m, &n)

	for m != 0 {
		temp := m
		m = n % m
		n = temp
	}

	fmt.Print(n)
}
