package main

import "fmt"

func main() {
	var m, n int

	fmt.Scan(&m, &n)

	max := n
	if m > n {
		max = m
	}
	lcm := max

	for lcm%n != 0 || lcm%m != 0 {
		lcm += max
	}

	fmt.Print(lcm)
}
