package main

import "fmt"

func main() {
	var i, poin, user, silverUser, platinumUser, goldUser int
	var stop bool

	silverUser = 0
	platinumUser = 0
	goldUser = 0

	fmt.Print("Masukkan jumlah user : ")
	fmt.Scan(&user)

	for i = 0; i < user; i++ {
		stop = false
		for !stop {
			fmt.Printf("Masukkan Poin user %d: ", i+1)
			fmt.Scan(&poin)

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
