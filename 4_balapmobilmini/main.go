package main

import (
	"fmt"
)

func main() {
	var car, winner string
	var aCount, bCount int

	fmt.Println("Masukkan urutan mobil (10 karakter, 'A' atau 'B'):")

	for i := 0; i < 10; i++ {
		fmt.Scan(&car)

		if car == "A" {
			aCount++
		} else if car == "B" {
			bCount++
		}

		if aCount == 4 && winner == "" {
			winner = "A"
		} else if bCount == 4 && winner == "" {
			winner = "B"
		}
	}

	if aCount != 5 || bCount != 5 {
		fmt.Println("tidak valid")
	} else {
		fmt.Println("Pemenang : ", winner)
	}
}
