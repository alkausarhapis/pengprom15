package main

import "fmt"

func main() {
	var currentVisitors, prevVisitors, totalVisitor, increase, increaseDays, days int
	var stop bool
	var averageIncrease float64

	fmt.Print("Visitors: ")
	fmt.Scan(&prevVisitors)
	stop = prevVisitors >= 0 && prevVisitors <= 200

	for stop {
		fmt.Print("Visitors: ")

		totalVisitor += prevVisitors
		days++

		fmt.Scan(&currentVisitors)

		increase = currentVisitors - prevVisitors
		if increase > 0 && currentVisitors <= 200 {
			increaseDays++
		}

		prevVisitors = currentVisitors
		stop = prevVisitors >= 0 && prevVisitors <= 200
	}

	if increaseDays > 0 {
		averageIncrease = float64(totalVisitor) / float64(days)
	}

	fmt.Printf("Jumlah hari naik: %d\n", increaseDays)
	fmt.Printf("Rata-rata kenaikan pengunjung per hari: %.2f\n", averageIncrease)
}
