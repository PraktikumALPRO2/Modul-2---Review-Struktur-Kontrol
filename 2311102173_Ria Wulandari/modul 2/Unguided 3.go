package main

import (
	"fmt"
)

func main() {
	var k float64

	// Meminta input nilai K
	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&k)

	// Menghitung nilai f(K) sesuai dengan persamaan
	atas := (4*k + 2) * (4*k + 2)
	bawah := (4*k + 1) * (4*k + 3)
	fk := atas / bawah

	// Menampilkan hasil perhitungan
	fmt.Printf("Nilai f(K) = %.10f\n", fk)
}
