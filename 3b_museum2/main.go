package main

import (
	"fmt"
)

func main() {
	var i, totalPengunjung, pengunjung, hari int
	var stop bool

	hari = 5

	for i = 1; i <= hari; i++ {
		stop = false
		for !stop {
			fmt.Printf("Hari %d: ", i)
			fmt.Scan(&pengunjung)

			stop = pengunjung >= 0 && pengunjung <= 200
		}

		totalPengunjung += pengunjung
	}

	fmt.Println("Jumlah pengunjung :", totalPengunjung, "orang")
}
