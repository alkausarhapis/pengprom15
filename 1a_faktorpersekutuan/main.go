package main

import "fmt"

func main() {
	var m, n, min, i int

	fmt.Scan(&m, &n)

	min = n
	if m < n {
		min = m
	}

	for i = 1; i <= min; i++ {
		if n%i == 0 && m%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
}
