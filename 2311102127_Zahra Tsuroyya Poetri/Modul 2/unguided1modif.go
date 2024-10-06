package main

import (
	"fmt"
)

func main() {
	var beratK1, beratK2 float64

	// Loop untuk meminta input berat belanjaan
	for {
		// Meminta input berat belanjaan untuk kedua kantong
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&beratK1, &beratK2)

		// Mengecek jika salah satu kantong berisi angka negatif atau total berat > 150kg
		if beratK1 < 0 || beratK2 < 0 || beratK1+beratK2 > 150 {
			break
		}

		// Menghitung selisih berat kedua kantong 
		var selisih float64
		if beratK1 > beratK2 {
			selisih = beratK1 - beratK2
		} else {
			selisih = beratK2 - beratK1
		}

		// Mengecek apakah selisih 9 kg atau lebih (motor pak Andi akan oleng)
		oleng := selisih >= 9

		// Menampilkan apakah motor oleng atau tidak
		fmt.Printf("Sepeda motor pak Andi akan oleng: %t\n", oleng)
	}

	// Menampilkan pesan proses selesai 
	fmt.Println("Proses selesai.")
}
