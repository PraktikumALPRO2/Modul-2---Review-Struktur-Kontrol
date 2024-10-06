// Agnes Refilina Fiska
// S1 IF11-05 / 2311102126

package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Print("Masukkan bilangan bulat (lebih dari 1): ")
	_, err := fmt.Scan(&number)
	if err != nil || number <= 1 {
		fmt.Println("Input tidak valid. Harap masukkan bilangan bulat lebih dari 1.")
		return
	}

	factors := []int{}
	isPrime := true

	// Mencari faktor
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			factors = append(factors, i) // Menambahkan faktor ke slice
			if i > 1 && i < number {
				isPrime = false // Menentukan jika ada faktor lain selain 1 dan number
			}
		}
	}

	// Menampilkan hasil
	fmt.Printf("Bilangan yang dimasukkan: %d\n", number)
	fmt.Print("Faktor: ")
	for idx, factor := range factors {
		if idx == len(factors)-1 {
			fmt.Printf("%d", factor) // Menampilkan faktor terakhir tanpa koma
		} else {
			fmt.Printf("%d, ", factor) // Menampilkan faktor dengan koma
		}
	}
	fmt.Println()

	// Menampilkan apakah bilangan tersebut prima
	if isPrime {
		fmt.Printf("%d adalah bilangan prima.\n", number)
	} else {
		fmt.Printf("%d bukan bilangan prima.\n", number)
	}
}
