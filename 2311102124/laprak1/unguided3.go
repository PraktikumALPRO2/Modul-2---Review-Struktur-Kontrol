/*Liya Khoirunnisa - 2311102124*/

package main

import (
	"fmt" // Untuk input atau output
	"math" // Untuk operasi matematika
)

// Fungsi untuk menghitung f(k)
func f(k float64) float64 {
	pembilang := math.Pow(4*k+2, 2)
	penyebut := (4*k + 1) * (4*k + 3)
	return pembilang / penyebut
}

func main() {
	var k float64
	fmt.Print("Nilai K = ") // Mencetak ke layar untuk meminta input
	fmt.Scanln(&k) // Membaca input
	hasil := f(k) // Memanggil fungsi
	fmt.Printf("Nilai f(K) = %.10f\n", hasil) // Mencetak nilai
}
