package main

import (
	"fmt"
)

func main() {
	var totalPengunjung, pengunjung int

	// Menerima input jumlah pengunjung selama 5 hari
	for i := 1; i <= 5; i++ {

		// Menerima input pengunjung pada hari ke-i
		for {
			fmt.Printf("Hari %d: ", i)
			fmt.Scan(&pengunjung)

			// Periksa apakah input valid
			if pengunjung >= 0 && pengunjung <= 200 {
				break
			}
		}

		// Tambahkan jumlah pengunjung ke total
		totalPengunjung += pengunjung
	}

	// Tampilkan jumlah total pengunjung selama 5 hari
	fmt.Printf("Jumlah pengunjung : %d", totalPengunjung)
}
