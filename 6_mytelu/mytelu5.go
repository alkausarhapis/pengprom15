package main

import "fmt"

func main() {
	var poin, totalPoin, silverUser, platinumUser, goldUser, totalSilver, totalGold, totalPlatinum int
	var averageSilver, averagePlatinum, averageGold float32
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
			totalSilver += poin
			silverUser++
		} else if poin >= 100 && poin <= 200 {
			totalPlatinum += poin
			platinumUser++
		} else {
			totalGold += poin
			goldUser++
		}
	}

	if silverUser > 0 {
		averageSilver = float32(totalSilver) / float32(silverUser)
	}

	if platinumUser > 0 {
		averagePlatinum = float32(totalPlatinum) / float32(platinumUser)
	}

	if goldUser > 0 {
		averageGold = float32(totalGold) / float32(goldUser)
	}

	fmt.Println("Average Silver User =", averageSilver)
	fmt.Println("Average Platinum User =", averagePlatinum)
	fmt.Println("Average Gold User =", averageGold)

}
