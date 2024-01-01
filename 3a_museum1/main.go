package main

import "fmt"

func main() {
	var currentVisitors, prevVisitors, totalVisitor, increase, increaseDays, days int

	var averageIncrease float64

	// Membaca jumlah pengunjung untuk hari pertama
	fmt.Print("Visitors: ")
	fmt.Scan(&prevVisitors)

	// Loop membaca jumlah pengunjung per hari
	for prevVisitors >= 0 && prevVisitors <= 200 {
		fmt.Print("Visitors: ")

		if currentVisitors > 200 || currentVisitors < 0 {
			break
		}

		// Menambahkan jumlah pengunjung ke totalVisitor
		totalVisitor += prevVisitors
		days++

		// Membaca jumlah pengunjung untuk hari berikutnya
		fmt.Scan(&currentVisitors)

		// Menghitung kenaikan pengunjung
		increase = currentVisitors - prevVisitors
		if increase > 0 && currentVisitors <= 200 {
			increaseDays++
		}

		prevVisitors = currentVisitors
	}

	if increaseDays > 0 {
		averageIncrease = float64(totalVisitor) / float64(days)
	}

	fmt.Printf("Jumlah hari naik: %d\n", increaseDays)
	fmt.Printf("Rata-rata kenaikan pengunjung per hari: %.2f\n", averageIncrease)
}
