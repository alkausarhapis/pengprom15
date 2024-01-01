package main

import "fmt"

func main() {
	var poin int

	for {
		fmt.Print("Masukkan Poin : ")
		fmt.Scan(&poin)

		if poin >= 50 {
			break
		}
	}

	if poin >= 50 && poin <= 99 {
		fmt.Print("Silver User")
	} else if poin >= 100 && poin <= 200 {
		fmt.Print("Platinum User")
	} else {
		fmt.Print("Gold User")
	}
}
