package main

import (
	"fmt"
)

func main() {
	var beratK1, beratK2 float64

	// Loop untuk meminta input hingga salah satu kantong berisi 9Kg atau lebih
	for {
		// Meminta input berat belanjaan untuk kedua kantong
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&beratK1, &beratK2)

		// Mengecek apakah salah satu kantong berisi 9Kg atau lebih
		if beratK1 >= 9 || beratK2 >= 9 {
			break
		}
	}

	// Menampilkan pesan proses selesai
	fmt.Println("Proses selesai")
}
