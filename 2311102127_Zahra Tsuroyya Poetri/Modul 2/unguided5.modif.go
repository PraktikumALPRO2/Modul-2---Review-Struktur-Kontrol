package main

import (
	"fmt"
)

func main() {
	var b int

	// Meminta input bilangan bulat b > 1
	fmt.Print("Bilangan: ")
	fmt.Scanln(&b)

	// Memastikan input valid
	if b <= 1 {
		fmt.Println("Bilangan harus lebih dari 1.")
		return
	}

	// Array untuk menyimpan faktor 
	var faktor [100]int
	hasilFaktor := 0

	// Mencari faktor bilangan
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			faktor[hasilFaktor] = i // Memasukkan faktor ke dalam array
			hasilFaktor++           // Menambah jumlah faktor
		}
	}

	// Menampilkan semua faktor
	fmt.Print("Faktor: ")
	for i := 0; i < hasilFaktor; i++ {
		fmt.Print(faktor[i], " ")
	}

	// Menentukan apakah bilangan tersebut prima
	prima := hasilFaktor == 2 

	// Menampilkan hasil apakah bilangan tersebut prima atau tidak
	fmt.Printf("\nPrima: %t\n", prima)
}
