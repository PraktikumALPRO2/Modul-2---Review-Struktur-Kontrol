package main

import (
	"fmt"
)

func main() {
	// Meminta input nilai K
	var k float64
	fmt.Print("Nilai K = ")
	fmt.Scanln(&k)

	// Menghitung bagian pembilang dan bagian penyebut dari persamaan
	pembilang := (4*k + 2) * (4*k + 2) // Menghitung (4k+2)^2
	penyebut := (4*k + 1) * (4*k + 3) // Menghitung (4k+1)(4k+3)
	fk := pembilang / penyebut // Menghitung f(K)

	// Menampilkan hasil nilai K 
	fmt.Printf("Nilai f(K) = %.10f\n", fk)
}
