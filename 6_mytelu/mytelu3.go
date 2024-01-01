package main

import "fmt"

func main() {
	var poin, user, silverUser, platinumUser, goldUser int

	silverUser = 0
	platinumUser = 0
	goldUser = 0

	fmt.Print("Masukkan jumlah user : ")
	fmt.Scan(&user)

	for i := 0; i < user; i++ {
		for {
			fmt.Printf("Masukkan Poin user %d: ", i+1)
			fmt.Scan(&poin)

			if poin >= 50 {
				break
			}
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
