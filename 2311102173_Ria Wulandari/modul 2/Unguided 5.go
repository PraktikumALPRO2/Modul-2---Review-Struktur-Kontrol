package main

import (
	"fmt"
)

func main() {
	// Input bilangan bulat
	var bilangan int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&bilangan)

	// Cek apakah bilangan > 1
	if bilangan <= 1 {
		fmt.Println("Bilangan harus lebih besar dari 1")
		return
	}

	// Temukan faktor-faktor bilangan
	fmt.Printf("Faktor: ")
	faktor := cariFaktor(bilangan)

	// Cek apakah bilangan prima
	isPrima := cekPrima(bilangan)

	// Tampilkan hasil faktor
	for _, f := range faktor {
		fmt.Printf("%d ", f)
	}

	// Tampilkan hasil apakah bilangan prima atau tidak
	fmt.Printf("\nPrima: %v\n", isPrima)
}

// Fungsi untuk mencari faktor-faktor bilangan
func cariFaktor(b int) []int {
	var faktor []int
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			faktor = append(faktor, i)
		}
	}
	return faktor
}

// Fungsi untuk mengecek apakah bilangan adalah bilangan prima
func cekPrima(b int) bool {
	if b < 2 {
		return false
	}
	for i := 2; i*i <= b; i++ {
		if b%i == 0 {
			return false
		}
	}
	return true
}
