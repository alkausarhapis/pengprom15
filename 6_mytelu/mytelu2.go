package main

import "fmt"

func main() {
	var poin int
	var stop bool

	stop = false

	for !stop {
		fmt.Print("Masukkan Poin : ")
		fmt.Scan(&poin)

		stop = poin >= 50
	}

	if poin >= 50 && poin <= 99 {
		fmt.Print("Silver User")
	} else if poin >= 100 && poin <= 200 {
		fmt.Print("Platinum User")
	} else {
		fmt.Print("Gold User")
	}
}
