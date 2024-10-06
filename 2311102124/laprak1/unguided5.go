/*Liya Khoirunnisa - 2311102124*/

package main

import (
	"fmt"
)

func main() {
	// Dekalrasi variabel
	var bilangan int
	fmt.Print("Bilangan: ") // Meminta input pengguna
	fmt.Scan(&bilangan) // Membaca input pengguna

	// Memeriksa bilangan
	if bilangan <= 1 {
		fmt.Println("Bilangan harus lebih besar dari 1.")
		return // Menghentikan program
	}

	fmt.Print("Faktor: ")
	jumlahFaktor := 0 // Variabel jumlah faktor
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Printf("%d ", i) // Menampilkan faktor
			jumlahFaktor++ // Menambah jumlah faktor
		}
	}
	fmt.Println()

	// Mengecek bilangan prima
	prima := jumlahFaktor == 2
	fmt.Printf("Prima: %v\n", prima)
}
