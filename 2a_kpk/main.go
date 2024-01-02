package main

import "fmt"

func main() {
	var m, n, kpk, max int

	fmt.Scan(&m, &n)

	max = n
	if m > n {
		max = m
	}
	kpk = max

	for kpk%n != 0 || kpk%m != 0 {
		kpk += max
	}

	fmt.Print(kpk)
}
