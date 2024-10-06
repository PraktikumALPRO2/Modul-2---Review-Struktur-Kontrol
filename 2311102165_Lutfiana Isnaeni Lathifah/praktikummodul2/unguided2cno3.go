// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
)

// Fungsi untuk mencari faktor-faktor dari sebuah bilangan
func cariFaktor(b int) []int {
	var faktor []int
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			faktor = append(faktor, i)
		}
	}
	return faktor
}

// Fungsi untuk menentukan apakah bilangan tersebut prima
func cekPrima(b int) bool {
	// Sebuah bilangan prima hanya memiliki 2 faktor: 1 dan dirinya sendiri
	if b <= 1 {
		return false
	}
	count := 0
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			count++
		}
	}
	return count == 2
}

func main() {
	// Input dari pengguna
	var b int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&b)

	// Cari faktor-faktor dari bilangan tersebut
	faktor := cariFaktor(b)
	fmt.Printf("Faktor: %v\n", faktor)

	// Cek apakah bilangan prima atau bukan
	prima := cekPrima(b)
	fmt.Printf("Prima: %v\n", prima)
}
