package main

import "fmt"

func main() {
	var poin, totalPoin, silverUser, platinumUser, goldUser int
	var stop bool

	silverUser = 0
	platinumUser = 0
	goldUser = 0
	i := 1
	totalPoin = 0

	for totalPoin < 500 {
		stop = false
		for !stop {
			fmt.Printf("Masukkan Poin user %d : ", i)
			fmt.Scan(&poin)
			totalPoin += poin

			if poin >= 50 {
				i++
			}

			stop = poin >= 50
		}

		if poin >= 50 && poin <= 99 {
			silverUser++
		} else if poin >= 100 && poin <= 200 {
			platinumUser++
		} else {
			goldUser++
		}
	}

	fmt.Println("Silver User =", silverUser)
	fmt.Println("Platinum User =", platinumUser)
	fmt.Println("Gold User =", goldUser)
}
