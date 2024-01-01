package main

import "fmt"

func main() {
	var m, n int

	fmt.Scan(&m, &n)

	for m != 0 {
		n, m = m, n%m
	}

	fmt.Print(n)
}
