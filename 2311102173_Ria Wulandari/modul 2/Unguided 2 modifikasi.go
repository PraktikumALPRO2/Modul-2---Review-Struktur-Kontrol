package main

import (
	"fmt"
	"math"
)

func main() {
	var berat1, berat2 float64

	for {
		// Meminta input berat masing-masing kantong
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		// Cek apakah berat salah satu kantong negatif
		if berat1 < 0 || berat2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		// Cek apakah total berat melebihi 150 kg
		if berat1+berat2 > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		// Cek apakah selisih berat lebih dari atau sama dengan 9 kg
		oleng := math.Abs(berat1-berat2) >= 9
		fmt.Println("Sepeda motor pak Andi akan oleng:", oleng)
	}
}
